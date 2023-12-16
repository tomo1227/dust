echo "pip Installing"
pip install --upgrade aws-sam-cli
sam --version

# localstack
# see: https://qiita.com/Brutus/items/89eb03be7c7bd9d6911d
echo "localstack Installing"
pip install localstack
localstack --version

pip install awscli-local
awslocal --version

go install github.com/cosmtrek/air@latest
go mod tidy
cd cmd/api
go mod tidy
air
