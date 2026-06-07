package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/model"
	"github.com/nanda/lenslift-backend/internal/repository"
)

type FoodService struct {
	foodRepo *repository.FoodRepository
	profileRepo *repository.ProfileRepository
}

func NewFoodService() *FoodService {
	return &FoodService{
		foodRepo: repository.NewFoodRepository(),
		profileRepo: repository.NewProfileRepository(),
	}
}

func (s *FoodService) AddFoodLog(userID string, req dto.FoodLogRequest) (*dto.FoodLogWithSummary, error) {
	userUUID, _ := uuid.Parse(userID)

	logDate := time.Now()
	if req.LogDate != "" {
		parsed, err := time.Parse("2006-01-02", req.LogDate)
		if err != nil {
			return nil, errors.New("format tanggal salah, gunakan YYYY-MM-DD")
		}
		logDate = parsed
	}

	food := &model.FoodLog{
		UserID: userUUID,
		FoodName: req.FoodName,
		Calories: req.Calories,
		ProteinG: req.ProteinG,
		CarbsG: req.CarbsG,
		FatG: req.FatG,
		Source: "manual",
		LogDate: time.Date(logDate.Year(), logDate.Month(), logDate.Day(), 0, 0, 0, 0, logDate.Location()),
	}

	if err := s.foodRepo.Create(food); err != nil {
		return nil, err
	}

	return s.buildFoodLogWithSummary(userID, food, logDate)
}

func (s *FoodService) AnalyzeFood(userID string, file multipart.File) (*dto.FoodLogWithSummary, error) {
	imageBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("gagal baca file")
	}

	base64Image := base64.StdEncoding.EncodeToString(imageBytes)

	aiResult, err := callClaudeVision(base64Image)
	if err != nil {
		return nil, err
	}

	userUUID, _ := uuid.Parse(userID)
	logDate := time.Now()

	food := &model.FoodLog{
		UserID: userUUID,
		FoodName: aiResult.FoodName,
		Calories: aiResult.Calories,
		ProteinG: aiResult.ProteinG,
		CarbsG: aiResult.CarbsG,
		FatG: aiResult.FatG,
		Source: "ai_photo",
		LogDate: time.Date(logDate.Year(), logDate.Month(), logDate.Day(), 0, 0, 0, 0, logDate.Location()),
	}

	if err := s.foodRepo.Create(food); err != nil {
		return nil, err
	}

	return s.buildFoodLogWithSummary(userID, food, logDate)
}

func (s *FoodService) GetDailyLogs(userID string, dateStr string) (*dto.DailyFoodLogs, error) {
	date := time.Now()
	if dateStr != "" {
		parsed, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, errors.New("format tanggal salah")
		}
		date = parsed
	}

	logs, err := s.foodRepo.FindByUserIDAndDate(userID, date)
	if err != nil {
		return nil, err
	}

	profile, _ := s.profileRepo.FindByUserID(userID)
	summary := s.calculateSummary(logs, profile)

	var logResponses []dto.FoodLogResponse
	for _, log := range logs {
		logResponses = append(logResponses, toFoodLogResponse(log))
	}

	if logResponses == nil {
		logResponses = []dto.FoodLogResponse{}
	}

	return &dto.DailyFoodLogs{
		Logs: logResponses,
		DailySummary: summary,
	}, nil
}

func (s *FoodService) DeleteFoodLog(userID string, foodID string) error {
	food, err := s.foodRepo.FindByID(foodID)
	if err != nil {
		return errors.New("log tidak ditemukan")
	}

	if food.UserID.String() != userID {
		return errors.New("akses ditolak")
	}

	return s.foodRepo.Delete(foodID)
}

func (s *FoodService) buildFoodLogWithSummary(userID string, food *model.FoodLog, date time.Time) (*dto.FoodLogWithSummary, error) {
	logs, err := s.foodRepo.FindByUserIDAndDate(userID, date)
	if err != nil {
		return nil, err
	}

	profile, _ := s.profileRepo.FindByUserID(userID)
	summary := s.calculateSummary(logs, profile)

	return &dto.FoodLogWithSummary{
		FoodLog: toFoodLogResponse(*food),
		DailySummary: summary,
	}, nil
}

func (s *FoodService) calculateSummary(logs []model.FoodLog, profile *model.Profile) dto.DailySummary {
	var totalCal, totalProtein, totalCarbs, totalFat float64
	for _, log := range logs {
		totalCal += log.Calories
		totalProtein += log.ProteinG
		totalCarbs += log.CarbsG
		totalFat += log.FatG
	}

	summary := dto.DailySummary{
		TotalCalories: totalCal,
		TotalProteinG: totalProtein,
		TotalCarbsG: totalCarbs,
		TotalFatG: totalFat,
	}

	if profile != nil {
		summary.TargetCalories = profile.TargetCalories
		summary.TargetProteinG = profile.TargetProteinG
	}

	return summary
}

func callClaudeVision(base64Image string) (*dto.AIAnalyzeResponse, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")

	requestBody := map[string]interface{}{
		"model": "claude-haiku-4-5",
		"max_tokens": 1024,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "image",
						"source": map[string]interface{}{
							"type": "base64",
							"media_type": "image/jpeg",
							"data": base64Image,
						},
					},
					{
						"type": "text",
						"text": `Kamu adalah analis nutrisi. Analisis makanan dalam foto ini dan jawab HANYA dengan JSON object, tanpa markdown, tanpa backtick. Format:
{"food_name": "nama makanan dalam bahasa Indonesia", "calories": 500, "protein_g": 25, "carbs_g": 60, "fat_g": 15}`,
					},
				},
			},
		},
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("gagal konek ke Claude API")
	}
	defer resp.Body.Close()

	var claudeResp struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&claudeResp); err != nil {
		return nil, errors.New("gagal parse response Claude")
	}

	if len(claudeResp.Content) == 0 {
		return nil, errors.New("response Claude kosong")
	}

	rawText := strings.TrimSpace(claudeResp.Content[0].Text)
	rawText = strings.ReplaceAll(rawText, "```json", "")
	rawText = strings.ReplaceAll(rawText, "```", "")
	rawText = strings.TrimSpace(rawText)
	
	var aiResult dto.AIAnalyzeResponse
	if err := json.Unmarshal([]byte(rawText), &aiResult); err != nil {
		return nil, fmt.Errorf("gagal parse hasil AI: %s", rawText)
	}

	return &aiResult, nil
}

func toFoodLogResponse(food model.FoodLog) dto.FoodLogResponse {
	return dto.FoodLogResponse{
		ID: food.ID.String(),
		FoodName: food.FoodName,
		Calories: food.Calories,
		ProteinG: food.ProteinG,
		CarbsG: food.CarbsG,
		FatG: food.FatG,
		Source: food.Source,
		LogDate: food.LogDate,
		CreatedAt: food.CreatedAt,
	}
}