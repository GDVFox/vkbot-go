package vkbotgo

// GENERATED FILE: DO NOT EDIT

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

// UsersGetParams params for users.get method
//
// https://vk.com/dev/users.get
type UsersGetParams struct {

	// User IDs or screen names (&#39;screen_name&#39;). By default, current user ID.
	UserIds []string `url:"user_ids,omitempty"`
	// Profile fields to return. Sample values: &#39;nickname&#39;, &#39;screen_name&#39;, &#39;sex&#39;, &#39;bdate&#39; (birthdate), &#39;city&#39;, &#39;country&#39;, &#39;timezone&#39;, &#39;photo&#39;, &#39;photo_medium&#39;, &#39;photo_big&#39;, &#39;has_mobile&#39;, &#39;contacts&#39;, &#39;education&#39;, &#39;online&#39;, &#39;counters&#39;, &#39;relation&#39;, &#39;last_seen&#39;, &#39;activity&#39;, &#39;can_write_private_message&#39;, &#39;can_see_all_posts&#39;, &#39;can_post&#39;, &#39;universities&#39;,
	Fields []string `url:"fields,omitempty"`
	// Case for declension of user name and surname: &#39;nom&#39; — nominative (default), &#39;gen&#39; — genitive , &#39;dat&#39; — dative, &#39;acc&#39; — accusative , &#39;ins&#39; — instrumental , &#39;abl&#39; — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *UsersGetParams) Validate() error {
	if len(param.UserIds) > 1000 {
		return fmt.Errorf("length(=%d) of parameter UserIds must be less or equal than 1000", len(param.UserIds))
	}

	enum2 := map[string]struct{}{"nom": {}, "gen": {}, "dat": {}, "acc": {}, "ins": {}, "abl": {}}
	if _, ok := enum2[param.NameCase]; !ok {
		return fmt.Errorf("parameter NameCase(=%v) expected in [nom gen dat acc ins abl]", param.NameCase)
	}

	return nil
}

// UsersGet calls VK API method users.get. Returns detailed information on users.
//
// https://vk.com/dev/users.get
func (vkBot *VkBot) UsersGet(params *UsersGetParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("users.get", vals)
}

// UsersSearchParams params for users.search method
//
// https://vk.com/dev/users.search
type UsersSearchParams struct {

	// Search query string (e.g., &#39;Vasya Babich&#39;).
	Q string `url:"q,omitempty"`
	// Sort order: &#39;1&#39; — by date registered, &#39;0&#39; — by rating
	Sort int64 `url:"sort,omitempty"`
	// Offset needed to return a specific subset of users.
	Offset int64 `url:"offset,omitempty"`
	// Number of users to return.
	Count int64 `url:"count,omitempty"`
	// Profile fields to return. Sample values: &#39;nickname&#39;, &#39;screen_name&#39;, &#39;sex&#39;, &#39;bdate&#39; (birthdate), &#39;city&#39;, &#39;country&#39;, &#39;timezone&#39;, &#39;photo&#39;, &#39;photo_medium&#39;, &#39;photo_big&#39;, &#39;has_mobile&#39;, &#39;rate&#39;, &#39;contacts&#39;, &#39;education&#39;, &#39;online&#39;,
	Fields []string `url:"fields,omitempty"`
	// City ID.
	City int64 `url:"city,omitempty"`
	// Country ID.
	Country int64 `url:"country,omitempty"`
	// City name in a string.
	Hometown string `url:"hometown,omitempty"`
	// ID of the country where the user graduated.
	UniversityCountry int64 `url:"university_country,omitempty"`
	// ID of the institution of higher education.
	University int64 `url:"university,omitempty"`
	// Year of graduation from an institution of higher education.
	UniversityYear int64 `url:"university_year,omitempty"`
	// Faculty ID.
	UniversityFaculty int64 `url:"university_faculty,omitempty"`
	// Chair ID.
	UniversityChair int64 `url:"university_chair,omitempty"`
	// &#39;1&#39; — female, &#39;2&#39; — male, &#39;0&#39; — any (default)
	Sex int64 `url:"sex,omitempty"`
	// Relationship status: &#39;1&#39; — Not married, &#39;2&#39; — In a relationship, &#39;3&#39; — Engaged, &#39;4&#39; — Married, &#39;5&#39; — It&#39;s complicated, &#39;6&#39; — Actively searching, &#39;7&#39; — In love
	Status int64 `url:"status,omitempty"`
	// Minimum age.
	AgeFrom int64 `url:"age_from,omitempty"`
	// Maximum age.
	AgeTo int64 `url:"age_to,omitempty"`
	// Day of birth.
	BirthDay int64 `url:"birth_day,omitempty"`
	// Month of birth.
	BirthMonth int64 `url:"birth_month,omitempty"`
	// Year of birth.
	BirthYear int64 `url:"birth_year,omitempty"`
	// &#39;1&#39; — online only, &#39;0&#39; — all users
	Online bool `url:"online,omitempty"`
	// &#39;1&#39; — with photo only, &#39;0&#39; — all users
	HasPhoto bool `url:"has_photo,omitempty"`
	// ID of the country where users finished school.
	SchoolCountry int64 `url:"school_country,omitempty"`
	// ID of the city where users finished school.
	SchoolCity int64 `url:"school_city,omitempty"`
	//
	SchoolClass int64 `url:"school_class,omitempty"`
	// ID of the school.
	School int64 `url:"school,omitempty"`
	// School graduation year.
	SchoolYear int64 `url:"school_year,omitempty"`
	// Users&#39; religious affiliation.
	Religion string `url:"religion,omitempty"`
	// Users&#39; interests.
	Interests string `url:"interests,omitempty"`
	// Name of the company where users work.
	Company string `url:"company,omitempty"`
	// Job position.
	Position string `url:"position,omitempty"`
	// ID of a community to search in communities.
	GroupID int64 `url:"group_id,omitempty"`
	//
	FromList []string `url:"from_list,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *UsersSearchParams) Validate() error {
	enum1 := map[int64]struct{}{0: {}, 1: {}}
	if _, ok := enum1[param.Sort]; !ok {
		return fmt.Errorf("parameter Sort(=%v) expected in [0 1]", param.Sort)
	}

	if param.Offset < 0.000000 {
		return fmt.Errorf("parameter Offset(=%v) must be greater or equal than 0.000000", param.Offset)
	}

	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.Count > 1000.000000 {
		return fmt.Errorf("parameter Count(=%v) must be less or equal than 1000.000000", param.Count)
	}

	if param.City < 0.000000 {
		return fmt.Errorf("parameter City(=%v) must be greater or equal than 0.000000", param.City)
	}

	if param.Country < 0.000000 {
		return fmt.Errorf("parameter Country(=%v) must be greater or equal than 0.000000", param.Country)
	}

	if param.UniversityCountry < 0.000000 {
		return fmt.Errorf("parameter UniversityCountry(=%v) must be greater or equal than 0.000000", param.UniversityCountry)
	}

	if param.University < 0.000000 {
		return fmt.Errorf("parameter University(=%v) must be greater or equal than 0.000000", param.University)
	}

	if param.UniversityYear < 0.000000 {
		return fmt.Errorf("parameter UniversityYear(=%v) must be greater or equal than 0.000000", param.UniversityYear)
	}

	if param.UniversityFaculty < 0.000000 {
		return fmt.Errorf("parameter UniversityFaculty(=%v) must be greater or equal than 0.000000", param.UniversityFaculty)
	}

	if param.UniversityChair < 0.000000 {
		return fmt.Errorf("parameter UniversityChair(=%v) must be greater or equal than 0.000000", param.UniversityChair)
	}

	if param.Sex < 0.000000 {
		return fmt.Errorf("parameter Sex(=%v) must be greater or equal than 0.000000", param.Sex)
	}

	enum13 := map[int64]struct{}{0: {}, 1: {}, 2: {}}
	if _, ok := enum13[param.Sex]; !ok {
		return fmt.Errorf("parameter Sex(=%v) expected in [0 1 2]", param.Sex)
	}

	if param.Status < 0.000000 {
		return fmt.Errorf("parameter Status(=%v) must be greater or equal than 0.000000", param.Status)
	}

	enum14 := map[int64]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}}
	if _, ok := enum14[param.Status]; !ok {
		return fmt.Errorf("parameter Status(=%v) expected in [0 1 2 3 4 5 6 7]", param.Status)
	}

	if param.AgeFrom < 0.000000 {
		return fmt.Errorf("parameter AgeFrom(=%v) must be greater or equal than 0.000000", param.AgeFrom)
	}

	if param.AgeTo < 0.000000 {
		return fmt.Errorf("parameter AgeTo(=%v) must be greater or equal than 0.000000", param.AgeTo)
	}

	if param.BirthDay < 0.000000 {
		return fmt.Errorf("parameter BirthDay(=%v) must be greater or equal than 0.000000", param.BirthDay)
	}

	if param.BirthMonth < 0.000000 {
		return fmt.Errorf("parameter BirthMonth(=%v) must be greater or equal than 0.000000", param.BirthMonth)
	}

	if param.BirthYear < 0.000000 {
		return fmt.Errorf("parameter BirthYear(=%v) must be greater or equal than 0.000000", param.BirthYear)
	}

	if param.SchoolCountry < 0.000000 {
		return fmt.Errorf("parameter SchoolCountry(=%v) must be greater or equal than 0.000000", param.SchoolCountry)
	}

	if param.SchoolCity < 0.000000 {
		return fmt.Errorf("parameter SchoolCity(=%v) must be greater or equal than 0.000000", param.SchoolCity)
	}

	if param.SchoolClass < 0.000000 {
		return fmt.Errorf("parameter SchoolClass(=%v) must be greater or equal than 0.000000", param.SchoolClass)
	}

	if param.School < 0.000000 {
		return fmt.Errorf("parameter School(=%v) must be greater or equal than 0.000000", param.School)
	}

	if param.SchoolYear < 0.000000 {
		return fmt.Errorf("parameter SchoolYear(=%v) must be greater or equal than 0.000000", param.SchoolYear)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// UsersSearch calls VK API method users.search. Returns a list of users matching the search criteria.
//
// https://vk.com/dev/users.search
func (vkBot *VkBot) UsersSearch(params *UsersSearchParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("users.search", vals)
}

// UsersIsAppUserParams params for users.isAppUser method
//
// https://vk.com/dev/users.isAppUser
type UsersIsAppUserParams struct {

	//
	UserID int64 `url:"user_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *UsersIsAppUserParams) Validate() error {
	return nil
}

// UsersIsAppUser calls VK API method users.isAppUser. Returns information whether a user installed the application.
//
// https://vk.com/dev/users.isAppUser
func (vkBot *VkBot) UsersIsAppUser(params *UsersIsAppUserParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("users.isAppUser", vals)
}

// UsersGetSubscriptionsParams params for users.getSubscriptions method
//
// https://vk.com/dev/users.getSubscriptions
type UsersGetSubscriptionsParams struct {

	// User ID.
	UserID int64 `url:"user_id,omitempty"`
	// &#39;1&#39; — to return a combined list of users and communities, &#39;0&#39; — to return separate lists of users and communities (default)
	Extended bool `url:"extended,omitempty"`
	// Offset needed to return a specific subset of subscriptions.
	Offset int64 `url:"offset,omitempty"`
	// Number of users and communities to return.
	Count int64 `url:"count,omitempty"`
	//
	Fields []string `url:"fields,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *UsersGetSubscriptionsParams) Validate() error {
	if param.UserID < 0.000000 {
		return fmt.Errorf("parameter UserID(=%v) must be greater or equal than 0.000000", param.UserID)
	}

	if param.Offset < 0.000000 {
		return fmt.Errorf("parameter Offset(=%v) must be greater or equal than 0.000000", param.Offset)
	}

	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.Count > 200.000000 {
		return fmt.Errorf("parameter Count(=%v) must be less or equal than 200.000000", param.Count)
	}

	return nil
}

// UsersGetSubscriptions calls VK API method users.getSubscriptions. Returns a list of IDs of users and communities followed by the user.
//
// https://vk.com/dev/users.getSubscriptions
func (vkBot *VkBot) UsersGetSubscriptions(params *UsersGetSubscriptionsParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("users.getSubscriptions", vals)
}

// UsersGetFollowersParams params for users.getFollowers method
//
// https://vk.com/dev/users.getFollowers
type UsersGetFollowersParams struct {

	// User ID.
	UserID int64 `url:"user_id,omitempty"`
	// Offset needed to return a specific subset of followers.
	Offset int64 `url:"offset,omitempty"`
	// Number of followers to return.
	Count int64 `url:"count,omitempty"`
	// Profile fields to return. Sample values: &#39;nickname&#39;, &#39;screen_name&#39;, &#39;sex&#39;, &#39;bdate&#39; (birthdate), &#39;city&#39;, &#39;country&#39;, &#39;timezone&#39;, &#39;photo&#39;, &#39;photo_medium&#39;, &#39;photo_big&#39;, &#39;has_mobile&#39;, &#39;rate&#39;, &#39;contacts&#39;, &#39;education&#39;, &#39;online&#39;.
	Fields []string `url:"fields,omitempty"`
	// Case for declension of user name and surname: &#39;nom&#39; — nominative (default), &#39;gen&#39; — genitive , &#39;dat&#39; — dative, &#39;acc&#39; — accusative , &#39;ins&#39; — instrumental , &#39;abl&#39; — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *UsersGetFollowersParams) Validate() error {
	if param.UserID < 0.000000 {
		return fmt.Errorf("parameter UserID(=%v) must be greater or equal than 0.000000", param.UserID)
	}

	if param.Offset < 0.000000 {
		return fmt.Errorf("parameter Offset(=%v) must be greater or equal than 0.000000", param.Offset)
	}

	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.Count > 1000.000000 {
		return fmt.Errorf("parameter Count(=%v) must be less or equal than 1000.000000", param.Count)
	}

	enum4 := map[string]struct{}{"nom": {}, "gen": {}, "dat": {}, "acc": {}, "ins": {}, "abl": {}}
	if _, ok := enum4[param.NameCase]; !ok {
		return fmt.Errorf("parameter NameCase(=%v) expected in [nom gen dat acc ins abl]", param.NameCase)
	}

	return nil
}

// UsersGetFollowers calls VK API method users.getFollowers. Returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
//
// https://vk.com/dev/users.getFollowers
func (vkBot *VkBot) UsersGetFollowers(params *UsersGetFollowersParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("users.getFollowers", vals)
}

// UsersReportParams params for users.report method
//
// https://vk.com/dev/users.report
type UsersReportParams struct {

	// ID of the user about whom a complaint is being made.
	UserID int64 `url:"user_id"`
	// Type of complaint: &#39;porn&#39; – pornography, &#39;spam&#39; – spamming, &#39;insult&#39; – abusive behavior, &#39;advertisment&#39; – disruptive advertisements
	Type string `url:"type"`
	// Comment describing the complaint.
	Comment string `url:"comment,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *UsersReportParams) Validate() error {
	if param.UserID < 0.000000 {
		return fmt.Errorf("parameter UserID(=%v) must be greater or equal than 0.000000", param.UserID)
	}

	enum1 := map[string]struct{}{"porn": {}, "spam": {}, "insult": {}, "advertisment": {}}
	if _, ok := enum1[param.Type]; !ok {
		return fmt.Errorf("parameter Type(=%v) expected in [porn spam insult advertisment]", param.Type)
	}

	return nil
}

// UsersReport calls VK API method users.report. Reports (submits a complain about) a user.
//
// https://vk.com/dev/users.report
func (vkBot *VkBot) UsersReport(params *UsersReportParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("users.report", vals)
}

// UsersGetNearbyParams params for users.getNearby method
//
// https://vk.com/dev/users.getNearby
type UsersGetNearbyParams struct {

	// geographic latitude of the place a user is located, in degrees (from -90 to 90)
	Latitude float64 `url:"latitude"`
	// geographic longitude of the place a user is located, in degrees (from -180 to 180)
	Longitude float64 `url:"longitude"`
	// current location accuracy in meters
	Accuracy int64 `url:"accuracy,omitempty"`
	// time when a user disappears from location search results, in seconds
	Timeout int64 `url:"timeout,omitempty"`
	// search zone radius type (1 to 4), :* 1 – 300 m,, :* 2 – 2400 m,, :* 3 – 18 km,, :* 4 – 150 km.
	Radius int64 `url:"radius,omitempty"`
	// list of additional fields to return. Available values: sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters, screen_name, maiden_name, timezone, occupation
	Fields []string `url:"fields,omitempty"`
	// Case for declension of user name and surname: , nom –nominative (default) , gen – genitive , dat – dative , acc – accusative , ins – instrumental , abl – prepositional
	NameCase string `url:"name_case,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *UsersGetNearbyParams) Validate() error {
	if param.Latitude < -90.000000 {
		return fmt.Errorf("parameter Latitude(=%v) must be greater or equal than -90.000000", param.Latitude)
	}

	if param.Latitude > 90.000000 {
		return fmt.Errorf("parameter Latitude(=%v) must be less or equal than 90.000000", param.Latitude)
	}

	if param.Longitude < -180.000000 {
		return fmt.Errorf("parameter Longitude(=%v) must be greater or equal than -180.000000", param.Longitude)
	}

	if param.Longitude > 180.000000 {
		return fmt.Errorf("parameter Longitude(=%v) must be less or equal than 180.000000", param.Longitude)
	}

	if param.Accuracy < 0.000000 {
		return fmt.Errorf("parameter Accuracy(=%v) must be greater or equal than 0.000000", param.Accuracy)
	}

	if param.Timeout < 0.000000 {
		return fmt.Errorf("parameter Timeout(=%v) must be greater or equal than 0.000000", param.Timeout)
	}

	if param.Radius < 0.000000 {
		return fmt.Errorf("parameter Radius(=%v) must be greater or equal than 0.000000", param.Radius)
	}

	enum4 := map[int64]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}}
	if _, ok := enum4[param.Radius]; !ok {
		return fmt.Errorf("parameter Radius(=%v) expected in [0 1 2 3 4]", param.Radius)
	}

	enum6 := map[string]struct{}{"nom": {}, "gen": {}, "dat": {}, "acc": {}, "ins": {}, "abl": {}}
	if _, ok := enum6[param.NameCase]; !ok {
		return fmt.Errorf("parameter NameCase(=%v) expected in [nom gen dat acc ins abl]", param.NameCase)
	}

	return nil
}

// UsersGetNearby calls VK API method users.getNearby. Indexes current user location and returns nearby users.
//
// https://vk.com/dev/users.getNearby
func (vkBot *VkBot) UsersGetNearby(params *UsersGetNearbyParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("users.getNearby", vals)
}
