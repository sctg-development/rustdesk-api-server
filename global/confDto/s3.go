package confDto

type S3Config struct {
	Region       string
	Endpoint     string
	AccessKey    string
	SecretKey    string
	Bucket       string
	Windows64Key string
	Windows32Key string
	OSXKey       string
	OSXArm64Key  string
	IOSKey       string
}
