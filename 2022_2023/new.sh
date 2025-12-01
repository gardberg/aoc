read -p "Day number: " day

mkdir "2023/"$day"a"
mkdir "2023/"$day"b"

cp "2023/main.go" "2023/"$day"a/main.go"

#curl https://adventofcode.com/2023/day/$day/input > "2023/"$day"a/input.txt"
#cp "2023/"$day"a/input.txt" "2023/"$day"b/input.txt"

touch "2023/"$day"a/input.txt"
touch "2023/"$day"a/test.txt"
