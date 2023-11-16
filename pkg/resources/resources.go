package resources

type Resource struct {
	Name       string
	Attributes map[string]string
}

var EC2 = Resource{
	Name:       "aws_instance",
	Attributes: map[string]string{"instance_type": "t2.micro"},
}

var Resources = map[string]Resource{
	"mxgraph.aws3.ec2": EC2,
}
