package utils

const (
	Base_url               = `https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/%[1]v?date=%[2]v`
	LayoutISO              = "2006-01-02"
	CommandDeleteRedisData = "docker exec redis redis-cli FLUSHALL"
)
