export CERTIFICATE_PATH=./jeremyrickard-github-io.pem
helm install ratify ratify/ratify \
    --atomic --namespace gatekeeper-system \
    --set cosign.enabled=false \
    --set-file notaryCert=${CERTIFICATE_PATH} \
     --set featureFlags.RATIFY_DYNAMIC_PLUGINS=true \
    --set image.tag=v1.0.0-rc.3