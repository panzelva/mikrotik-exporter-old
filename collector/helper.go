package collector

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

func metricStringCleanup(in string) string {
	return strings.Replace(in, "-", "_", -1)
}

func descriptionForPropertyName(prefix, property string, labelNames []string) *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, prefix, metricStringCleanup(property)),
		property,
		labelNames,
		nil,
	)
}

func description(prefix, name, helpText string, labelNames []string) *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, prefix, name),
		helpText,
		labelNames,
		nil,
	)
}

func splitStringToFloats(metric string) (float64, float64, error) {
	strs := strings.Split(metric, ",")
	if len(strs) == 0 {
		return 0, 0, nil
	}
	m1, err := strconv.ParseFloat(strs[0], 64)
	if err != nil {
		return math.NaN(), math.NaN(), err
	}
	m2, err := strconv.ParseFloat(strs[1], 64)
	if err != nil {
		return math.NaN(), math.NaN(), err
	}
	return m1, m2, nil
}

func parseStringToFloat64(value, device, property, errorMsg string) (float64, error) {
	parsedValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		if errorMsg != "" {
			log.WithFields(log.Fields{"device": device, "property": property, "value": value, "error": err}).Error(errorMsg)
		}
		return math.NaN(), err
	}

	return parsedValue, nil
}

func parseUptime(uptime string) (float64, error) {
	var u time.Duration

	reMatch := uptimeRegex.FindAllStringSubmatch(uptime, -1)

	// should get one and only one match back on the regex
	if len(reMatch) != 1 {
		return 0, fmt.Errorf("invalid uptime value sent to regex")
	}

	for i, match := range reMatch[0] {
		if match != "" && i != 0 {
			v, err := strconv.Atoi(match)
			if err != nil {
				log.WithFields(log.Fields{
					"uptime": uptime,
					"value":  match,
					"error":  err,
				}).Error("error parsing uptime field value")
				return float64(0), err
			}
			u += time.Duration(v) * uptimeParts[i-1]
		}
	}

	return u.Seconds(), nil
}
