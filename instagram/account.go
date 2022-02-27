//
// account.go
// Copyright 2017 Konstantin Dovnar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// http://www.apache.org/licenses/LICENSE-2.0
//

package instagram

import "encoding/json"

// An Account describes an Instagram account info.
type Account struct {
	Biography       string
	ExternalURL     string
	Followers       uint32
	Follows         uint32
	FullName        string
	ID              string
	Private         bool
	Verified        bool
	MediaCount      uint32
	ProfilePicURL   string
	ProfilePicURLhd string
	Username        string
}

// Update try to update account info
func (a *Account) Update() error {
	account, err := GetAccountByUsername(a.Username)
	if err != nil {
		return err
	}
	*a = account
	return nil
}

func getFromAccountPage(data []byte) (Account, bool) {
	var accountJSON struct {
		SeoCategoryInfos      [][]string `json:"seo_category_infos"`
		LoggingPageID         string     `json:"logging_page_id"`
		ShowSuggestedProfiles bool       `json:"show_suggested_profiles"`
		ShowFollowDialog      bool       `json:"show_follow_dialog"`
		Graphql               struct {
			User struct {
				Biography              string      `json:"biography"`
				BlockedByViewer        bool        `json:"blocked_by_viewer"`
				RestrictedByViewer     interface{} `json:"restricted_by_viewer"`
				CountryBlock           bool        `json:"country_block"`
				ExternalURL            string      `json:"external_url"`
				ExternalURLLinkshimmed string      `json:"external_url_linkshimmed"`
				EdgeFollowedBy         struct {
					Count int `json:"count"`
				} `json:"edge_followed_by"`
				Fbid             string `json:"fbid"`
				FollowedByViewer bool   `json:"followed_by_viewer"`
				EdgeFollow       struct {
					Count int `json:"count"`
				} `json:"edge_follow"`
				FollowsViewer         bool        `json:"follows_viewer"`
				FullName              string      `json:"full_name"`
				HasArEffects          bool        `json:"has_ar_effects"`
				HasClips              bool        `json:"has_clips"`
				HasGuides             bool        `json:"has_guides"`
				HasChannel            bool        `json:"has_channel"`
				HasBlockedViewer      bool        `json:"has_blocked_viewer"`
				HighlightReelCount    int         `json:"highlight_reel_count"`
				HasRequestedViewer    bool        `json:"has_requested_viewer"`
				HideLikeAndViewCounts bool        `json:"hide_like_and_view_counts"`
				ID                    string      `json:"id"`
				IsBusinessAccount     bool        `json:"is_business_account"`
				IsProfessionalAccount bool        `json:"is_professional_account"`
				IsEmbedsDisabled      bool        `json:"is_embeds_disabled"`
				IsJoinedRecently      bool        `json:"is_joined_recently"`
				BusinessAddressJSON   interface{} `json:"business_address_json"`
				BusinessContactMethod string      `json:"business_contact_method"`
				BusinessEmail         interface{} `json:"business_email"`
				BusinessPhoneNumber   interface{} `json:"business_phone_number"`
				BusinessCategoryName  interface{} `json:"business_category_name"`
				OverallCategoryName   interface{} `json:"overall_category_name"`
				CategoryEnum          interface{} `json:"category_enum"`
				CategoryName          interface{} `json:"category_name"`
				IsPrivate             bool        `json:"is_private"`
				IsVerified            bool        `json:"is_verified"`
				EdgeMutualFollowedBy  struct {
					Count int           `json:"count"`
					Edges []interface{} `json:"edges"`
				} `json:"edge_mutual_followed_by"`
				ProfilePicURL            string        `json:"profile_pic_url"`
				ProfilePicURLHd          string        `json:"profile_pic_url_hd"`
				RequestedByViewer        bool          `json:"requested_by_viewer"`
				ShouldShowCategory       bool          `json:"should_show_category"`
				ShouldShowPublicContacts bool          `json:"should_show_public_contacts"`
				Username                 string        `json:"username"`
				ConnectedFbPage          interface{}   `json:"connected_fb_page"`
				Pronouns                 []interface{} `json:"pronouns"`
				EdgeFelixVideoTimeline   struct {
					Count    int `json:"count"`
					PageInfo struct {
						HasNextPage bool        `json:"has_next_page"`
						EndCursor   interface{} `json:"end_cursor"`
					} `json:"page_info"`
					Edges []interface{} `json:"edges"`
				} `json:"edge_felix_video_timeline"`
				EdgeOwnerToTimelineMedia struct {
					Count    int `json:"count"`
					PageInfo struct {
						HasNextPage bool   `json:"has_next_page"`
						EndCursor   string `json:"end_cursor"`
					} `json:"page_info"`
					Edges []struct {
						Node struct {
							Typename   string `json:"__typename"`
							ID         string `json:"id"`
							Shortcode  string `json:"shortcode"`
							Dimensions struct {
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"dimensions"`
							DisplayURL            string `json:"display_url"`
							EdgeMediaToTaggedUser struct {
								Edges []struct {
									Node struct {
										User struct {
											FullName         string `json:"full_name"`
											FollowedByViewer bool   `json:"followed_by_viewer"`
											ID               string `json:"id"`
											IsVerified       bool   `json:"is_verified"`
											ProfilePicURL    string `json:"profile_pic_url"`
											Username         string `json:"username"`
										} `json:"user"`
										X float64 `json:"x"`
										Y float64 `json:"y"`
									} `json:"node"`
								} `json:"edges"`
							} `json:"edge_media_to_tagged_user"`
							FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
							FactCheckInformation   interface{} `json:"fact_check_information"`
							GatingInfo             interface{} `json:"gating_info"`
							SharingFrictionInfo    struct {
								ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
								BloksAppURL               interface{} `json:"bloks_app_url"`
							} `json:"sharing_friction_info"`
							MediaOverlayInfo interface{} `json:"media_overlay_info"`
							MediaPreview     interface{} `json:"media_preview"`
							Owner            struct {
								ID       string `json:"id"`
								Username string `json:"username"`
							} `json:"owner"`
							IsVideo              bool   `json:"is_video"`
							HasUpcomingEvent     bool   `json:"has_upcoming_event"`
							AccessibilityCaption string `json:"accessibility_caption"`
							DashInfo             struct {
								IsDashEligible    bool        `json:"is_dash_eligible"`
								VideoDashManifest interface{} `json:"video_dash_manifest"`
								NumberOfQualities int         `json:"number_of_qualities"`
							} `json:"dash_info"`
							HasAudio           bool   `json:"has_audio"`
							TrackingToken      string `json:"tracking_token"`
							VideoURL           string `json:"video_url"`
							VideoViewCount     int    `json:"video_view_count"`
							EdgeMediaToCaption struct {
								Edges []struct {
									Node struct {
										Text string `json:"text"`
									} `json:"node"`
								} `json:"edges"`
							} `json:"edge_media_to_caption"`
							EdgeMediaToComment struct {
								Count int `json:"count"`
							} `json:"edge_media_to_comment"`
							CommentsDisabled bool `json:"comments_disabled"`
							TakenAtTimestamp int  `json:"taken_at_timestamp"`
							EdgeLikedBy      struct {
								Count int `json:"count"`
							} `json:"edge_liked_by"`
							EdgeMediaPreviewLike struct {
								Count int `json:"count"`
							} `json:"edge_media_preview_like"`
							Location struct {
								ID            string `json:"id"`
								HasPublicPage bool   `json:"has_public_page"`
								Name          string `json:"name"`
								Slug          string `json:"slug"`
							} `json:"location"`
							ThumbnailSrc       string `json:"thumbnail_src"`
							ThumbnailResources []struct {
								Src          string `json:"src"`
								ConfigWidth  int    `json:"config_width"`
								ConfigHeight int    `json:"config_height"`
							} `json:"thumbnail_resources"`
							FelixProfileGridCrop      interface{}   `json:"felix_profile_grid_crop"`
							CoauthorProducers         []interface{} `json:"coauthor_producers"`
							ProductType               string        `json:"product_type"`
							ClipsMusicAttributionInfo interface{}   `json:"clips_music_attribution_info"`
							EdgeSidecarToChildren     struct {
								Edges []struct {
									Node struct {
										Typename   string `json:"__typename"`
										ID         string `json:"id"`
										Shortcode  string `json:"shortcode"`
										Dimensions struct {
											Height int `json:"height"`
											Width  int `json:"width"`
										} `json:"dimensions"`
										DisplayURL            string `json:"display_url"`
										EdgeMediaToTaggedUser struct {
											Edges []struct {
												Node struct {
													User struct {
														FullName         string `json:"full_name"`
														FollowedByViewer bool   `json:"followed_by_viewer"`
														ID               string `json:"id"`
														IsVerified       bool   `json:"is_verified"`
														ProfilePicURL    string `json:"profile_pic_url"`
														Username         string `json:"username"`
													} `json:"user"`
													X float64 `json:"x"`
													Y float64 `json:"y"`
												} `json:"node"`
											} `json:"edges"`
										} `json:"edge_media_to_tagged_user"`
										FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
										FactCheckInformation   interface{} `json:"fact_check_information"`
										GatingInfo             interface{} `json:"gating_info"`
										SharingFrictionInfo    struct {
											ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
											BloksAppURL               interface{} `json:"bloks_app_url"`
										} `json:"sharing_friction_info"`
										MediaOverlayInfo interface{} `json:"media_overlay_info"`
										MediaPreview     string      `json:"media_preview"`
										Owner            struct {
											ID       string `json:"id"`
											Username string `json:"username"`
										} `json:"owner"`
										IsVideo              bool   `json:"is_video"`
										HasUpcomingEvent     bool   `json:"has_upcoming_event"`
										AccessibilityCaption string `json:"accessibility_caption"`
									} `json:"node"`
								} `json:"edges"`
							} `json:"edge_sidecar_to_children"`
						} `json:"node,omitempty"`
					} `json:"edges"`
				} `json:"edge_owner_to_timeline_media"`
				EdgeSavedMedia struct {
					Count    int `json:"count"`
					PageInfo struct {
						HasNextPage bool        `json:"has_next_page"`
						EndCursor   interface{} `json:"end_cursor"`
					} `json:"page_info"`
					Edges []interface{} `json:"edges"`
				} `json:"edge_saved_media"`
				EdgeMediaCollections struct {
					Count    int `json:"count"`
					PageInfo struct {
						HasNextPage bool        `json:"has_next_page"`
						EndCursor   interface{} `json:"end_cursor"`
					} `json:"page_info"`
					Edges []interface{} `json:"edges"`
				} `json:"edge_media_collections"`
				EdgeRelatedProfiles struct {
					Edges []interface{} `json:"edges"`
				} `json:"edge_related_profiles"`
			} `json:"user"`
		} `json:"graphql"`
		ToastContentOnLoad      interface{} `json:"toast_content_on_load"`
		ShowViewShop            bool        `json:"show_view_shop"`
		ProfilePicEditSyncProps struct {
			ShowChangeProfilePicConfirmDialog      bool        `json:"show_change_profile_pic_confirm_dialog"`
			ShowProfilePicSyncReminders            bool        `json:"show_profile_pic_sync_reminders"`
			IdentityID                             string      `json:"identity_id"`
			RemoveProfilePicHeader                 interface{} `json:"remove_profile_pic_header"`
			RemoveProfilePicSubtext                interface{} `json:"remove_profile_pic_subtext"`
			RemoveProfilePicConfirmCta             interface{} `json:"remove_profile_pic_confirm_cta"`
			RemoveProfilePicCancelCta              interface{} `json:"remove_profile_pic_cancel_cta"`
			IsBusinessCentralIdentity              bool        `json:"is_business_central_identity"`
			ChangeProfilePicActionsScreenHeader    []string    `json:"change_profile_pic_actions_screen_header"`
			ChangeProfilePicActionsScreenSubheader interface{} `json:"change_profile_pic_actions_screen_subheader"`
			ChangeProfilePicActionsScreenUploadCta []string    `json:"change_profile_pic_actions_screen_upload_cta"`
			ChangeProfilePicActionsScreenRemoveCta []string    `json:"change_profile_pic_actions_screen_remove_cta"`
			ChangeProfilePicActionsScreenCancelCta []string    `json:"change_profile_pic_actions_screen_cancel_cta"`
			ChangeProfilePicHeader                 interface{} `json:"change_profile_pic_header"`
			ChangeProfilePicSubtext                interface{} `json:"change_profile_pic_subtext"`
		} `json:"profile_pic_edit_sync_props"`
		AlwaysShowMessageButtonToProAccount bool `json:"always_show_message_button_to_pro_account"`
	}

	err := json.Unmarshal(data, &accountJSON)
	if err != nil {
		return Account{}, false
	}

	account := Account{}
	account.Biography = accountJSON.Graphql.User.Biography
	account.ExternalURL = accountJSON.Graphql.User.ExternalURL
	account.FullName = accountJSON.Graphql.User.FullName
	account.ID = accountJSON.Graphql.User.ID
	account.Private = accountJSON.Graphql.User.IsPrivate
	account.Verified = accountJSON.Graphql.User.IsVerified
	account.ProfilePicURL = accountJSON.Graphql.User.ProfilePicURL
	account.ProfilePicURLhd = accountJSON.Graphql.User.ProfilePicURLHd
	account.Username = accountJSON.Graphql.User.Username
	account.Followers = uint32(accountJSON.Graphql.User.EdgeFollowedBy.Count)
	account.Follows = uint32(accountJSON.Graphql.User.EdgeFollow.Count)
	account.MediaCount = uint32(accountJSON.Graphql.User.EdgeMediaCollections.Count)

	return account, true
}

func getFromSearchPage(data []byte) ([]Account, error) {
	var searchJSON struct {
		Users []struct {
			Position int `json:"position"`
			User     struct {
				Pk            string `json:"pk"`
				Username      string `json:"username"`
				FullName      string `json:"full_name"`
				IsPrivate     bool   `json:"is_private"`
				ProfilePicURL string `json:"profile_pic_url"`
				IsVerified    bool   `json:"is_verified"`
				FollowerCount int    `json:"follower_count"`
			} `json:"user"`
		} `json:"users"`
	}

	err := json.Unmarshal(data, &searchJSON)
	if err != nil {
		return nil, err
	}

	accounts := []Account{}

	for _, user := range searchJSON.Users {
		account := Account{}
		account.ID = user.User.Pk
		account.Username = user.User.Username
		account.FullName = user.User.FullName
		account.Private = user.User.IsPrivate
		account.Verified = user.User.IsVerified
		account.ProfilePicURL = user.User.ProfilePicURL
		account.Followers = uint32(user.User.FollowerCount)

		accounts = append(accounts, account)
	}

	return accounts, nil
}
