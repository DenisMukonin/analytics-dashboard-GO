package models

import (
	"fmt",
	"time",
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnalyticsEvent struct {
	ID					primitive.ObjectID			`bson:"_id,omitempty" json:"id"`
	UserID			string									`bson:"userId" json:"userId"`
	EventType   string                  `bson:"eventType" json:"eventType"` // page_view, click, purchase
	Properties  map[string]interface{}  `bson:"properties" json:"properties"`
	Timestamp   time.Time               `bson:"timestamp" json:"timestamp"`
}

type Metric struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Date        time.Time           `bson:"date" json:"date"`
	MetricName  string              `bson:"metricName" json:"metricName"` // users, revenue, conversions
	Value       float64             `bson:"value" json:"value"`
	Trend       float64             `bson:"trend,omitempty" json:"trend,omitempty"`
	Unit        string              `bson:"unit" json:"unit"` // "users", "$", "%"
	Dimensions  map[string]string   `bson:"dimensions,omitempty" json:"dimensions,omitempty"` // region, platform
}

type Widget struct {
	ID        string                  `bson:"id" json:"id"`
	Type      string                  `bson:"type" json:"type"` // line-chart, bar-chart, metric
	Title     string                  `bson:"title" json:"title"`
	Config    map[string]interface{}  `bson:"config" json:"config"`
	Position  Position                `bson:"position" json:"position"`
	Size      Size                    `bson:"size" json:"size"`
}

type Position struct {
	X  int  `bson:"x" json:"x"`
	Y  int  `bson:"y" json:"y"`
}

type Size struct {
	W  int  `bson:"w" json:"w"`
	H  int  `bson:"h" json:"h"`
}

type DashboardConfig struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	UserID     string              `bson:"userId" json:"userId"`
	Name       string              `bson:"name" json:"name"`
	Widgets    []Widget            `bson:"widgets" json:"widgets"`
	CreatedAt  time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time           `bson:"updatedAt" json:"updatedAt"`
}

type ApiResponse struct {
	Data     interface{}  `json:"data"`
	Status   int          `json:"status"`
	Message  string       `json:"message,omitempty"`
}

type TimeRangeRequest struct {
	StartDate  string    `form:"startDate" binding:"required"`
	EndDate    string    `form:"endDate" binding:"required"`
	Metrics    []string  `form:"metrics"`
}