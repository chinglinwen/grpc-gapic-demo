#!/bin/bash
# still draft

find generated/cli/ -name '*.go' | xargs sed -i 's,"google.golang.org/protobuf/jsonpb",jsonpb "google.golang.org/protobuf/encoding/protojson",g'

matchText="err = jsonpb.Unmarshal(in"
text="$(
	cat <<'EOF'
            b, err :=io.ReadAll(in)
			if err!=nil{
				return err
			}
			
			err = jsonpb.Unmarshal(b
EOF
)"
echo "$text"

find generated/cli/ -name '*.go' | xargs sed -i sed "s/$matchText=\"err = jsonpb.Unmarshal(in)\"/$matchText=\"$text\"/"
