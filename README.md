<p align="center">
  <img src="https://cdn.evntaly.com/Resources/og.png" alt="Evntaly Cover" width="100%">
</p>

<h3 align="center">Evntaly</h3>

<p align="center">
 An advanced event tracking and analytics platform designed to help developers capture, analyze, and react to user interactions efficiently.
</p>
<p align="center">
  <a href="https://pkg.go.dev/github.com/Evntaly/evntaly-go"><img src="https://img.shields.io/github/v/tag/Evntaly/evntaly-go" alt="Go Version"></a>
  <a href="https://github.com/Evntaly/evntaly-go"><img src="https://img.shields.io/github/license/Evntaly/evntaly-go" alt="license"></a>
</p>

# evntaly-go

**EvntalySDK** is a Go client for interacting with the Evntaly event tracking platform. It provides methods to initialize tracking, log events, identify users, and check API usage limits.

## Features

- **Initialize** the SDK with a developer secret and project token.
- **Track events** with metadata and tags.
- **Identify users** for personalization and analytics.
- **Enable or disable** tracking globally.

## Installation

To install the SDK using `go get`:

```sh
 go get github.com/Evntaly/evntaly-go
```

## Usage

### Initialization

Initialize the SDK with your developer secret and project token:

```go
package main

import (
    "github.com/Evntaly/evntaly-go"
    "fmt"
)

func main() {
    evntaly := evntaly.NewEvntalySDK("YOUR_DEVELOPER_SECRET", "YOUR_PROJECT_TOKEN")
    evntaly.SetRequestTimeout(5 * time.Second) // optional

    fmt.Println("Evntaly SDK initialized!")
}
```

### Tracking Events

To track an event:

```go
 event := evntaly.Event{
     Title:       "Payment Received",
     Description: "User completed a purchase",
     Message:     "Order #12345",
     Data: map[string]interface{}{
         "user_id": "67890",
         "timestamp": "2025-01-08T09:30:00Z",
         "referrer": "social_media",
         "email_verified": true,
         "amount": 149.99,
         "currency": "USD",
         "payment_method": "Credit Card",
     },
     Tags:       []string{"purchase", "payment", "ecommerce"},
     Notify:     true,
     Icon:       "💰",
     ApplyRuleOnly: false,
     User:       evntaly.EventUser{ID: "12345"},
     Type:       "Transaction",
     SessionID:  "20750ebc-dabf-4fd4-9498-443bf30d6095_bsd",
     Feature:    "Checkout",
     Topic:      "@Sales",
 }
 err := evntaly.Track(event)
 if err != nil {
     fmt.Println("Error tracking event:", err)
 }
```

### Identifying Users

To identify a user:

```go
 user := evntaly.User{
     ID:           "12345",
     Email:        "user@example.com",
     FullName:     "John Doe",
     Organization: "ExampleCorp",
     Data: map[string]interface{}{
         "id": "JohnD",
         "email": "user@example.com",
         "location": "USA",
         "salary": 75000,
         "timezone": "America/New_York",
         "subscription_plan": "Premium",
         "last_login": "2025-02-24T15:30:00Z",
     },
 }
 err := evntaly.IdentifyUser(user)
 if err != nil {
     fmt.Println("Error identifying user:", err)
 }
```

### Enabling/Disabling Tracking

Control event tracking globally:

```go
 evntaly.DisableTracking()  // Disables tracking
 evntaly.EnableTracking()   // Enables tracking
```

## License

This project is licensed under the MIT License.

---

*Note: Replace **`"YOUR_DEVELOPER_SECRET"`** and **`"YOUR_PROJECT_TOKEN"`** with actual credentials.*