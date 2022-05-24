output_path="pack/auto3mad"

#cd backend && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../$output_path/backend/backend && cd ..
cd backend && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../$output_path/backend/backend.exe && cd ..
cp -r backend/conf $output_path/backend/
cp -r backend/sql $output_path/backend/

mkdir -p $output_path/frontend
cp -r frontend/dist $output_path/frontend/

tar -C pack -czvf auto3mad.tar.gz auto3mad
