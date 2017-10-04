///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	image "gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/client/image"
	models "gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/models"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/vscli/i18n"
)

var (
	getImagesLong = i18n.T(`Get images.`)

	// TODO: add examples
	getImagesExample = i18n.T(``)
)

// NewCmdGetImage creates command responsible for getting images.
func NewCmdGetImage(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "image [IMAGE]",
		Short:   i18n.T("Get images"),
		Long:    getImagesLong,
		Example: getImagesExample,
		Args:    cobra.MaximumNArgs(1),
		Aliases: []string{"images"},
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			if len(args) > 0 {
				err = getImage(out, errOut, cmd, args)
			} else {
				err = getImages(out, errOut, cmd)
			}
			CheckErr(err)
		},
	}
	return cmd
}

func getImage(out, errOut io.Writer, cmd *cobra.Command, args []string) error {
	client := imageManagerClient()
	params := &image.GetImageByNameParams{
		Context:   context.Background(),
		ImageName: args[0],
	}
	resp, err := client.Image.GetImageByName(params)
	if err != nil {
		fmt.Println("list images returned an error")
		return err
	}
	return formatImageOutput(out, false, []*models.Image{resp.Payload})
}

func getImages(out, errOut io.Writer, cmd *cobra.Command) error {
	client := imageManagerClient()
	params := &image.GetImagesParams{
		Context: context.Background(),
	}
	resp, err := client.Image.GetImages(params)
	if err != nil {
		fmt.Println("list images returned an error")
		return err
	}
	return formatImageOutput(out, true, resp.Payload)
}

func formatImageOutput(out io.Writer, list bool, images []*models.Image) error {
	if vsConfig.Json {
		encoder := json.NewEncoder(out)
		encoder.SetIndent("", "    ")
		if list {
			return encoder.Encode(images)
		}
		return encoder.Encode(images[0])
	}
	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Name", "URL", "BaseImage", "Status", "Created Date"})
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetCenterSeparator("")
	for _, image := range images {
		table.Append([]string{*image.Name, image.DockerURL, *image.BaseImageName, string(image.Status), time.Unix(image.CreatedTime, 0).Local().Format(time.UnixDate)})
	}
	table.Render()
	return nil
}
