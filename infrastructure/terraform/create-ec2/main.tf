provider "aws" {
  region = "ap-southeast-1"
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-amd64-server-*"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "app_server" {
  count = 2
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.small"
  key_name = "devops"

  vpc_security_group_ids = [
    aws_security_group.sg_web.id
  ]

  tags = {
    Name = "terramform-ec2-${count.index + 1}"
  }
}

resource "aws_security_group" "sg_web" { 
    name = "terramform-ec2-sg-web"

    ingress  {
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
    ingress  {
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
    ingress  {
        from_port = 443
        to_port = 443
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}

output "pubblic_ip" {
    value = aws_instance.app_server[*].public_ip
}


