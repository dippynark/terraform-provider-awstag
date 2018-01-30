# Terraform AWSTag Provider

This is a simple provider which adds an AWS EC2 tag resource. This resource allows us to do the following:
- Modify tags for EC2 resources created in an external Terraform environment
- Create tags with dynamic keys (tag keys currently have to be specified as static strings)

More details about the need for this resource can be found in [this](https://github.com/terraform-providers/terraform-provider-aws/issues/3143) issue.

## Example

```
provider "awstag" {
  region = "eu-west-1"
}

provider "aws" {
  region = "eu-west-1"
}

resource "aws_subnet" "example" {
  vpc_id                  = "vpc-12345678"
  cidr_block              = "192.168.1.0/24"  

  lifecycle {
    ignore_changes = ["tags"]
  }

  tags {
    Name = "Example"
  }
}

variable "tag_key" {
  default = "my_tag_key"
}

variable "tag_value" {
  default = "my_tag_value"
}

resource "awstag_ec2_tag" "my_tag" {
  ec2_id = "${aws_subnet.example.id}"
  key = "${var.tag_key}"
  value = "${var.tag_value}"
}
```