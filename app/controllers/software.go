package controllers

import (
	"rustdesk-api-server/utils/beegoHelper"

	"context"

	"rustdesk-api-server/global"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/beego/beego/v2/server/web"
)

type SoftwareController struct {
	web.Controller
}

// Get software information
func (ctl *SoftwareController) GetSoftwareInfo() {
	ctl.Ctx.Output.JSON(beegoHelper.H{
		"software": "SCTGDesk",
		"version":  "1.2.4",
		"website":  "https://github.com/sctg-development/sctgdesk",
	}, true, false)

}

func getSignedReleaseUrl(bucket string, key string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(global.ConfigVar.S3.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(global.ConfigVar.S3.AccessKey, global.ConfigVar.S3.SecretKey, "")),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: global.ConfigVar.S3.Endpoint}, nil
			})))

	if err != nil {
		return "", err
	}
	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	resp, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(options *s3.PresignOptions) {})
	return resp.URL, err
}

// Get Windows client download link
func (ctl *SoftwareController) GetClientDownloadLink(key string) {

	url, err := getSignedReleaseUrl(global.ConfigVar.S3.Bucket, key)

	if err != nil {
		ctl.Ctx.Output.JSON(beegoHelper.H{
			"msg": "Failed to generate download link",
			"err": err.Error(),
		}, true, false)
		return
	}
	ctl.Ctx.Output.JSON(beegoHelper.H{
		"url": url,
	}, true, false)
}

func (ctl *SoftwareController) GetClientDownloadLinkW64() {
	ctl.GetClientDownloadLink(global.ConfigVar.S3.Windows64Key)
}

func (ctl *SoftwareController) GetClientDownloadLinkW32() {
	ctl.GetClientDownloadLink(global.ConfigVar.S3.Windows32Key)
}

func (ctl *SoftwareController) GetClientDownloadLinkOSX() {
	ctl.GetClientDownloadLink(global.ConfigVar.S3.OSXKey)
}

func (ctl *SoftwareController) GetClientDownloadLinkOSXArm64() {
	ctl.GetClientDownloadLink(global.ConfigVar.S3.OSXArm64Key)
}
