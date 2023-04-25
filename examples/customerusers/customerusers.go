package main

import (
	"flag"
	"fmt"

	"github.com/woodstock-tokyo/sendbird"
)

const (
	IssueAccessToken   = false
	IdleCustomType     = "idle"
	OccupiedCustomType = "occupied"
	GroupChannelType   = "group_channels"

	UserRoleMetaKey   = "role"
	UserRoleMetaValue = "customer"
)

var (
	apiKey   = flag.String("key", "", "API Key for using Sendbird Platform API")
	userID   = flag.String("id", "", "UserID for admin user registeration")
	nickName = flag.String("name", "", "Nick name for admin user registeration")
	//profileURL = flag.String("profile", "", "The URL of the user’s profile image.")
)

func main() {
	flag.Parse()

	testClient, err := sendbird.NewClient(sendbird.WithAPIKey(*apiKey))
	check(err)

	user, err := testClient.CreateAUserWithURL(&sendbird.CreateAUserWithURLRequest{
		UserID:   *userID,
		NickName: *nickName,
		//ProfileURL:       *profileURL,
		IssueAccessToken: IssueAccessToken,
	})
	check(err)
	fmt.Printf("User: %+v \n", user)

	meta := make(map[string]string)
	meta[UserRoleMetaKey] = UserRoleMetaValue

	returnMeta, err := testClient.CreateAnUserMetaData(user.UserID, &sendbird.CreateAnUserMetaDataRequest{
		MetaData: meta,
	})
	check(err)
	fmt.Printf("Usermeta: %+v \n", returnMeta)

	chResp, err := testClient.ListMyGroupChannels(user.UserID, &sendbird.ListMyGroupChannelsRequest{})
	check(err)
	fmt.Printf("Length of user's groupChannels: %d \n", len(chResp.Channels))

	if len(chResp.Channels) == 0 {
		groupChResp, err := testClient.ListGroupChannels(&sendbird.ListGroupChannelsRequest{
			CustomType: IdleCustomType,
		})
		check(err)
		fmt.Printf("Length of idle groupChannels: %d \n", len(groupChResp.Channels))

		if len(groupChResp.Channels) > 0 {
			targetCh := groupChResp.Channels[0]

			targetCh, err = testClient.UpdateAGroupChannel(targetCh.ChannelURL, &sendbird.UpdateAGroupChannelRequest{
				CustomType: OccupiedCustomType,
			})
			check(err)

			targetCh, err := testClient.InviteMembersToGroupChannel(targetCh.ChannelURL, &sendbird.InviteMembersToGroupChannelRequest{
				UserIDs: []string{user.UserID},
			})
			check(err)

			fmt.Printf("GroupChannel: %+v \n", targetCh)
		}

	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
