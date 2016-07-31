cid="replace_with_client_id"
cs="replace_with_client_secret"
sub="yoursubdomain"

#puts "https://#{sub}.slack.com/oauth/authorize?client_id=#{cid}&scope=users:write,chat:write:user,channels:read,channels:history,channels:write,groups:read,groups:write,groups:history,im:history,im:read,im:write,team:read,usergroups:read,usergroups:write,users:read"

code = "123"  # replace with code from ^

str = `curl "https://#{sub}.slack.com/api/oauth.access?client_id=#{cid}&client_secret=#{cs}&code=#{code}"`
puts str # will output token you need
