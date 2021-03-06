##############################################
# mapの使い方と使い所
##############################################

# 配列.map { |i| 動作させたい中身 }

# まず実際の動きを見てください

p [1, 3, 5].map { |i| i + 2 }
# => [3, 5, 7]

# 動作を一つづつ分解すると
# { |i| i + 2} に1が入力値として入って結果が配列の１つ目に詰められます
# [3]
# 次に3が同様に入力値として入ってきて配列の２つ目に詰められます
# [3, 5]
# 5も同様です
# [3, 5, 7]

# iに詰めていくことで複数行で書くこともできます

hoge = [1, 3, 5].map do |i|
  i = i + 2
  i + 2
end
p hoge

# mapの使い所は返却値が配列なことから配列を元にして新たに配列を作りたいときはmap
# そうでない場合は eachを使えばいいと思います
