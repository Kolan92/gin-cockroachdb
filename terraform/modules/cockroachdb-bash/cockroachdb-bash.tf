resource "null_resource" "setup_cocroachdb" {

  triggers = {
    timestamp = timestamp() #TODO replace with more relible way of forcing execution
  }

  provisioner "local-exec" {
    command     = "./${path.module}/setup_cocroachdb.sh"
    interpreter = ["bash", "-c"]
  }
}
