package evntaly

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// EvntalySDK is the main struct for the Evntaly SDK.
type EvntalySDK struct {
	BaseURL         string
	DeveloperSecret string
	ProjectToken    string
	TrackingEnabled bool
}

type EventUser struct {
	ID string `json:"id"`
}

type Event struct {
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
	Tags          []string    `json:"tags"`
	Notify        bool        `json:"notify"`
	Icon          string      `json:"icon"`
	ApplyRuleOnly bool        `json:"apply_rule_only"`
	User          EventUser   `json:"user"`
	Type          string      `json:"type"`
	SessionID     string      `json:"sessionID"`
	Feature       string      `json:"feature"`
	Topic         string      `json:"topic"`
}

type User struct {
	ID           string                 `json:"id"`
	Email        string                 `json:"email"`
	FullName     string                 `json:"full_name"`
	Organization string                 `json:"organization"`
	Data         map[string]interface{} `json:"data"`
}

func NewEvntalySDK(developerSecret, projectToken string) *EvntalySDK {
	return &EvntalySDK{
		BaseURL:         "https://evntaly.com/prod",
		DeveloperSecret: developerSecret,
		ProjectToken:    projectToken,
		TrackingEnabled: true,
	}
}

func (sdk *EvntalySDK) CheckLimit() (bool, error) {
	url := fmt.Sprintf("%s/api/v1/account/check-limits/%s", sdk.BaseURL, sdk.DeveloperSecret)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result map[string]bool
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, err
	}

	limitReached, exists := result["limitReached"]
	if !exists {
		return false, errors.New("unexpected API response format")
	}

	return !limitReached, nil
}

func (sdk *EvntalySDK) Track(event Event) error {
	if !sdk.TrackingEnabled {
		fmt.Println("Tracking is disabled. Event not sent.")
		return nil
	}

	canTrack, err := sdk.CheckLimit()
	if err != nil || !canTrack {
		fmt.Println("checkLimit returned false. Event not sent.")
		return err
	}

	url := fmt.Sprintf("%s/api/v1/register/event", sdk.BaseURL)
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(eventJSON))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("secret", sdk.DeveloperSecret)
	req.Header.Set("pat", sdk.ProjectToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to track event: status code %d", resp.StatusCode)
	}

	fmt.Println("✅ Event tracked successfully")
	return nil
}

func (sdk *EvntalySDK) IdentifyUser(user User) error {
	url := fmt.Sprintf("%s/api/v1/register/user", sdk.BaseURL)
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(userJSON))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("secret", sdk.DeveloperSecret)
	req.Header.Set("pat", sdk.ProjectToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to identify user: status code %d", resp.StatusCode)
	}

	fmt.Println("✅ User identified successfully")
	return nil
}

func (sdk *EvntalySDK) DisableTracking() {
	sdk.TrackingEnabled = false
	fmt.Println("🚫 Tracking disabled.")
}

func (sdk *EvntalySDK) EnableTracking() {
	sdk.TrackingEnabled = true
	fmt.Println("🟢 Tracking enabled.")
}
