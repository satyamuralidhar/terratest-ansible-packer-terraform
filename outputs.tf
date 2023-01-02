output "public_ip" {
  value = module.virtualmachine.public_ip
}
output "vm_resource_id" {
  value = module.virtualmachine.vm_resource_id
}

output "vnet_cidr" {
  value = module.virtualmachine.vnet_address_space
}

output "rsg" {
  value = module.virtualmachine.resource_group_name
}
output "location" {
  value = module.virtualmachine.location
}
output "packerimage" {
  value = module.virtualmachine.packer_image_name
}