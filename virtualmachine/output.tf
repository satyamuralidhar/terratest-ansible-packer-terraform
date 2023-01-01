output "public_ip" {
  value = azurerm_linux_virtual_machine.myvm.public_ip_address
}
output "vm_resource_id" {
  value = azurerm_linux_virtual_machine.myvm.id
}

output "vnet_address_space" {
  value = azurerm_virtual_network.myvnet.address_space
}

output "resource_group_name" {
  value = azurerm_virtual_network.myvnet.resource_group_name
}
output "location" {
  value = azurerm_virtual_network.myvnet.location
}
output "packer_image_name" {
  value = data.azurerm_image.packerimage.name
}