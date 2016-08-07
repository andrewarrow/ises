cid="replace_with_client_id"
cs="replace_with_client_secret"

['team1', 'team2', 'team3', 'team4'].each do |sub|
  puts "https://#{sub}.slack.com/oauth/authorize?client_id=#{cid}&scope=client"
end

code = "123"  # replace with code from ^

sub="team1"
#str = `curl "https://#{sub}.slack.com/api/oauth.access?client_id=#{cid}&client_secret=#{cs}&code=#{code}"`
#puts str # will output token you need
