package service

import (
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/model"
	"github.com/nanda/lenslift-backend/internal/repository"
)

type SessionService struct {
	sessionRepo *repository.SessionRepository
}

func NewSessionService() *SessionService {
	return &SessionService{
		sessionRepo: repository.NewSessionRepository(),
	}
}

func (s *SessionService) StartSession(userID string, req dto.StartSessionRequest) (*dto.SessionResponse, error) {
	userUUID, _ := uuid.Parse(userID)

	session := &model.WorkoutSession{
		UserID: userUUID,
	}

	if req.TemplateID != "" {
		templateUUID, err := uuid.Parse(req.TemplateID)
		if err != nil {
			return nil, errors.New("template id tidak valid")
		}
		session.TemplateID = &templateUUID
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}

	resp := &dto.SessionResponse{
		ID: session.ID.String(),
		StartedAt: session.StartedAt,
		EndedAt: session.EndedAt,
	}
	if session.TemplateID != nil {
		resp.TemplateID = session.TemplateID.String()
	}

	return resp, nil
}

func (s *SessionService) LogSet(userID string, sessionID string, req dto.LogSetRequest) (*dto.SessionLogResponse, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return nil, errors.New("sesi tidak ditemukan")
	}

	if session.UserID.String() != userID {
		return nil, errors.New("akses ditolak")
	}

	if session.EndedAt != nil {
		return nil, errors.New("sesi sudah selesai, tidak bisa log set")
	}

	sessionUUID, _ := uuid.Parse(sessionID)
	log := &model.SessionLog{
		SessionID: sessionUUID,
		ExerciseName: req.ExerciseName,
		SetNumber: req.SetNumber,
		ActualReps: req.ActualReps,
		ActualWeightKg: req.ActualWeightKg,
	}

	if err := s.sessionRepo.CreateLog(log); err != nil {
		return nil, err
	}

	return &dto.SessionLogResponse{
		ID: log.ID.String(),
		ExerciseName: log.ExerciseName,
		SetNumber: log.SetNumber,
		ActualReps: log.ActualReps,
		ActualWeightKg: log.ActualWeightKg,
		LoggedAt: log.LoggedAt,
	}, nil
}

func (s *SessionService) EndSession(userID string, sessionID string) (*dto.SessionDetailResponse, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return nil, errors.New("sesi tidak ditemukan")
	}

	if session.UserID.String() != userID {
		return nil, errors.New("akses ditolak")
	}

	if session.EndedAt != nil {
		return nil, errors.New("sesi sudah selesai")
	}

	now := time.Now()
	session.EndedAt = &now
	if err := s.sessionRepo.Update(session); err != nil {
		return nil, err
	}

	return buildSessionDetail(session), nil
}

func (s *SessionService) GetAllSessions(userID string) ([]dto.SessionResponse, error) {
	sessions, err := s.sessionRepo.FindAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	var resp []dto.SessionResponse
	for _, sess := range sessions {
		r := dto.SessionResponse{
			ID: sess.ID.String(),
			StartedAt: sess.StartedAt,
			EndedAt: sess.EndedAt,
		}
		if sess.TemplateID != nil {
			r.TemplateID = sess.TemplateID.String()
		}
		resp = append(resp, r)
	}
	return resp, nil
}

func (s *SessionService) GetSession(userID string, sessionID string) (*dto.SessionDetailResponse, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return nil, errors.New("sesi tidak ditemukan")
	}

	if session.UserID.String() != userID {
		return nil, errors.New("akses ditolak")
	}

	return buildSessionDetail(session), nil
}

func buildSessionDetail(session *model.WorkoutSession) *dto.SessionDetailResponse {
	var totalSets int
	var totalVolume float64

	logs := make([]dto.SessionLogResponse, len(session.Logs))
	for i, log := range session.Logs {
		logs[i] = dto.SessionLogResponse{
			ID: log.ID.String(),
			ExerciseName: log.ExerciseName,
			SetNumber: log.SetNumber,
			ActualReps: log.ActualReps,
			ActualWeightKg: log.ActualWeightKg,
			LoggedAt: log.LoggedAt,
		}
		totalSets++
		totalVolume += float64(log.ActualReps) * log.ActualWeightKg
	}

	resp := &dto.SessionDetailResponse{
		ID: session.ID.String(),
		StartedAt: session.StartedAt,
		EndedAt: session.EndedAt,
		TotalSets: totalSets,
		TotalVolumeKg: math.Round(totalVolume*100) / 100,
		Logs: logs,
	}

	if session.TemplateID != nil {
		resp.TemplateID = session.TemplateID.String()
	}

	if session.EndedAt != nil {
		duration := session.EndedAt.Sub(session.StartedAt).Minutes()
		resp.DurationMinutes = int(math.Round(duration))
	}

	return resp
}