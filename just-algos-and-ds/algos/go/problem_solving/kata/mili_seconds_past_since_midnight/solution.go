package mili_seconds_past_since_midnight

const MsPerSec = 1000
const MsPerMin = MsPerSec * 60
const MsPerHour = MsPerMin * 60

func solution(h, m, s int) int {
	return h*MsPerHour + m*MsPerMin + s*MsPerSec
}
