package gubconfig

import (
	"github.com/UltimateTournament/ArcadeGubernator/gubernator"
)

func GenGlobalIPConfig(uniqueKey string) *gubernator.RateLimitReq {
	return &gubernator.RateLimitReq{
		Name:      "global_ip",
		UniqueKey: uniqueKey,
		Hits:      1,
		Limit:     300, // 300/min
		Duration:  60000,
		Algorithm: 0,
		Behavior:  0,
	}
}

func GenSMSConfig(uniqueKey string) *gubernator.RateLimitReq {
	return &gubernator.RateLimitReq{
		Name:      "sms",
		UniqueKey: uniqueKey,
		Hits:      1,
		Limit:     1, // 1/30s
		Duration:  30000,
		Algorithm: 0,
		Behavior:  0,
	}
}

func GenRegisterConfig(uniqueKey string) *gubernator.RateLimitReq {
	return &gubernator.RateLimitReq{
		Name:      "register_ip",
		UniqueKey: uniqueKey,
		Hits:      1,
		Limit:     10, // 10/min
		Duration:  60000,
		Algorithm: 0,
		Behavior:  0,
	}
}
