package sun

import (
	"math"
	"time"
)

const (
	secondsInADay      = 86400
	unixEpochJulianDay = 2440587.5

	// Degree provides a precise fraction for converting between degrees and
	// radians.
	Degree = math.Pi / 180

	// J2000 is the Julian date for January 1, 2000, 12:00:00 TT.
	J2000 = 2451545
)

// TimeToJulianDay converts a time.Time into a Julian day.
func TimeToJulianDay(t time.Time) float64 {
	return float64(t.Unix())/secondsInADay + unixEpochJulianDay
}

// JulianDayToTime converts a Julian day into a time.Time.
func JulianDayToTime(d float64) time.Time {
	return time.Unix(int64((d-unixEpochJulianDay)*secondsInADay), 0).UTC()
}

// MeanSolarNoon calculates the time at which the sun is at its highest
// altitude. The returned time is in Julian days.
func MeanSolarNoon(longitude float64, year int, month time.Month, day int) float64 {
	t := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
	return TimeToJulianDay(t) - longitude/360
}

// SolarMeanAnomaly calculates the angle of the sun relative to the earth for
// the specified Julian day.
func SolarMeanAnomaly(d float64) float64 {
	v := math.Remainder(357.5291+0.98560028*(d-J2000), 360)
	if v < 0 {
		v += 360
	}
	return v
}

// EquationOfCenter calculates the angular difference between the position of
// the earth in its elliptical orbit and the position it would occupy in a
// circular orbit for the given mean anomaly.
func EquationOfCenter(solarAnomaly float64) float64 {
	var (
		anomalyInRad = solarAnomaly * (math.Pi / 180)
		anomalySin   = math.Sin(anomalyInRad)
		anomaly2Sin  = math.Sin(2 * anomalyInRad)
		anomaly3Sin  = math.Sin(3 * anomalyInRad)
	)
	return 1.9148*anomalySin + 0.0200*anomaly2Sin + 0.0003*anomaly3Sin
}

// ArgumentOfPerihelion calculates the argument of periapsis for the earth on
// the given Julian day.
func ArgumentOfPerihelion(d float64) float64 {
	return 102.93005 + 0.3179526*(d-2451545)/36525
}

// EclipticLongitude calculates the angular distance of the earth along the
// ecliptic.
func EclipticLongitude(solarAnomaly, equationOfCenter, d float64) float64 {
	return math.Mod(solarAnomaly+equationOfCenter+180+ArgumentOfPerihelion(d), 360)
}

// SolarTransit calculates the Julian data for the local true solar transit.
func SolarTransit(d, solarAnomaly, eclipticLongitude float64) float64 {
	equationOfTime := 0.0053*math.Sin(solarAnomaly*Degree) -
		0.0069*math.Sin(2*eclipticLongitude*Degree)
	return d + equationOfTime
}

// Declination calculates one of the two angles required to locate a point on
// the celestial sphere in the equatorial coordinate system. The ecliptic
// longitude parameter must be in degrees.
func Declination(eclipticLongitude float64) float64 {
	return math.Asin(math.Sin(eclipticLongitude*Degree)*0.39779) / Degree
}

// HourAngle calculates the second of the two angles required to locate a point
// on the celestial sphere in the equatorial coordinate system.
func HourAngle(latitude, declination float64) float64 {
	var (
		latitudeRad    = latitude * Degree
		declinationRad = declination * Degree
		numerator      = -0.01449 - math.Sin(latitudeRad)*math.Sin(declinationRad)
		denominator    = math.Cos(latitudeRad) * math.Cos(declinationRad)
	)

	// Check for no sunrise/sunset
	if numerator/denominator > 1 {
		// Sun never rises
		return math.MaxFloat64
	}

	if numerator/denominator < -1 {
		// Sun never sets
		return -1 * math.MaxFloat64
	}

	return math.Acos(numerator/denominator) / Degree
}

func SunriseSunset(latitude, longitude float64, date time.Time) (time.Time, time.Time) {
	var (
		d                 = MeanSolarNoon(longitude, date.Year(), date.Month(), date.Day())
		solarAnomaly      = SolarMeanAnomaly(d)
		equationOfCenter  = EquationOfCenter(solarAnomaly)
		eclipticLongitude = EclipticLongitude(solarAnomaly, equationOfCenter, d)
		solarTransit      = SolarTransit(d, solarAnomaly, eclipticLongitude)
		declination       = Declination(eclipticLongitude)
		hourAngle         = HourAngle(latitude, declination)
		frac              = hourAngle / 360
		sunrise           = solarTransit - frac
		sunset            = solarTransit + frac
	)

	// Check for no sunrise, no sunset
	if hourAngle == math.MaxFloat64 || hourAngle == -1*math.MaxFloat64 {
		return time.Time{}, time.Time{}
	}

	return JulianDayToTime(sunrise), JulianDayToTime(sunset)
}
