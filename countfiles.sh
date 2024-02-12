count=$(ls -R | sed '/^$/d' | sed '/^[.]/d' | wc -l)
let "count++"
echo $count
# | wc -l