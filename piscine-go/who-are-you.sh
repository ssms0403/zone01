returnValue=`curl -s https://learn.zone01dakar.sn/assets/superhero/all.json`
echo $returnValue  | cat > list.json
jq '.[]  | select(.id==70) | .name' list.json