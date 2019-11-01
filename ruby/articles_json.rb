require 'json'

##############################################
# jsonをファイルから読み込む場合
##############################################

# JSON.loadを使うとjsonがHashクラスに展開されます

f = File.open('assets/json_sample.json')
hoge = JSON.load(f)
puts hoge
puts hoge.class
f.close
