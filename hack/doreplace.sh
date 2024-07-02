doreplace () {
  old="$1"
  new="$2"
  find ./ -name '*.go' -type f -not -path './.git/*' -exec sed -i -e "s|${old}|${new}|g" {} \;
}