storage "consul" {
  address = "consul.example.com:8500"
  token   = "consul-acl-token-abcdef1234567890abcdef1234567890"
}

# HashiCorp Vault AppRole credentials
role_id    = "db02de05-fa39-4855-059b-67221c5c2f63"
secret_id  = "6a174c20-f6de-a53c-74d2-6018fcceff64"

# Nomad ACL token
NOMAD_TOKEN=00000000-0000-0000-0000-abcdefABCDEF
