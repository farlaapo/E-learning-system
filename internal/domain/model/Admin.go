package model


import (
	
)


type DashboardMetrics struct {
	TotalUsers    int64   `json:"total_users"`
	TotalCourses  int64   `json:"total_courses"`
	TotalRevenue  float64 `json:"total_revenue"`
	ActiveCourses int64   `json:"active_courses"`
}