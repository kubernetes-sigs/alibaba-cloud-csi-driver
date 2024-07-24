package interfaces

type NasClientFactoryInterface interface {
	V1(region string) (NasV1Interface, error)
	V2(region string) (NasClientV2Interface, error)
}
