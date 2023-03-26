VERSION=0.0.1

echo "Building badger-db:$VERSION..."

docker build . -t badger-db:$VERSION && cd ..

echo "Docker image created: badger-db:$VERSION"

docker tag badger-db:$VERSION icebaker/badger-db:$VERSION

docker push icebaker/badger-db:$VERSION

echo "Docker image pushed: badger-db:$VERSION"
