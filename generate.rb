cid="replace_with_client_id"
cs="replace_with_client_secret"

#puts "https://#{sub}.slack.com/oauth/authorize?client_id=#{cid}&scope=users:write,chat:write:user,channels:read,channels:history,channels:write,groups:read,groups:write,groups:history,im:history,im:read,im:write,team:read,usergroups:read,usergroups:write,users:read"
['team1', 'team2', 'team3', 'team4'].each do |sub|
  puts "https://#{sub}.slack.com/oauth/authorize?client_id=#{cid}&scope=client"
end

code = "123"  # replace with code from ^

sub="team1"
#str = `curl "https://#{sub}.slack.com/api/oauth.access?client_id=#{cid}&client_secret=#{cs}&code=#{code}"`
#puts str # will output token you need
