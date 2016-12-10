package secret

const newSecretTemplate = `apiVersion: v1
kind: Secret
metadata:
  name: %s
type: Opaque
data:
%s
`
