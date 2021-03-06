##############################################
# each文を使った繰り返し
##############################################

# each文は配列やハッシュに対して繰り返し処理が行なえます

# 配列の場合の繰り返し

[1, 3, 5].each do |r|
  p r
end

# 1
# 3
# 5

# ハッシュの場合の繰り返し

{a: 1, b: 2, c: 3}.each do |r|
  p r
end

# [:a, 1]
# [:b, 2]
# [:c, 3]

# 何番目の取り出しかを知りたい場合

# each_with_indexを使うことで iに 0から順番が入ります

[1, 3, 5].each_with_index do |r, i|
  p r
  p "#{i}番目"
end

# 1
# "0番目"
# 3
# "1番目"
# 5
# "2番目"
