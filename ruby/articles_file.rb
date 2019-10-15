##############################################
# ファイルを開いて中身を見る
##############################################

# ファイルを開くにはopenを使います
# ファイルクラスに展開されているのがわかります

p open('./assets/file1.txt')
# => #<File:./assets/file1.txt>

# ファイルを開いて一行づつ表示する
# ファイルの中身は
#
# aiueo
# kakikukeko
# sasisuseso

open('./assets/file1.txt').each do |l|
  p l
end

# => "aiueo\n"
#    "kakikukeko\n"
#    "sasisuseso\n"

# ファイルをメモリ上に一度に読み込む場合
# 一度メモリ上に展開してから動作させるので高速に動作しますが、
# あまりに巨大なファイルの場合メモリの上限で落ちる可能性もあります

f1 = open('./assets/file1.txt')
f1.each do |l|
  p l
end
f1.close

##############################################
# ファイルの存在を確認する
##############################################

# File.exist?('ファイルの置き場所')

# ファイルが存在すればtrueをそうでない場合falseを返します

p File.exist?('./assets/file1.txt')
# => true

p File.exist?('./assets/hoge.txt')
# => false

##############################################
# 一時ファイルを作成する
##############################################

# 一時ファイルの作成にはTempfileを使います。

# 一時ファイルは以下の条件で削除されてしまうので気をつけましょう。
# 1. closeが呼び出された時
# 2. GCが呼び出された時
# 3. Rubyのインタプリタ終了時（プロセス終了時）

# ブロックを出るときにcloseが呼び出されるので
# Tempfileで作成されたものはでた時点でなくなっています
# 1のcloseの条件にあてはまります
require 'tempfile'
Tempfile.create('hoge') do |t|
  p t.path
  # => "/var/folders/00/3815xdpn6wjfhb5r082_xpfr0000gp/T/hoge20191012-60237-1iqvivb"
end

require 'tempfile'
hogePath = ''
Tempfile.create('hoge') do |t|
  hogePath = t.path
  File.exist?(hogePath)
  # => true
end

File.exist?(hogePath)
# => false

##############################################
# ファイルに書き込み
##############################################

# ファイルを新規で作成して書き込む場合（既にファイルがあった場合中身が破棄されます）
# 第2引数にwを指定します

f2 = File.open('./assets/write_test1.txt', 'w')
f2.puts 'hoga2'
f2.close

# ファイルに追記して書き込む場合
# 第2引数にaを指定します

f3 = File.open('./assets/write_test2.txt', 'a')
f3.puts 'hoga2'
f3.close

##############################################
# ファイルをコピーするやり方(Fileutils.cp)
##############################################

# File.openを使ってやるやり方もありますが、Fileutilsを使うのが楽です。

# FileUtils.cp(コピー元ファイル,コピー先ファイル)

FileUtils.cp('./assets/file1.txt', './assets/copied_file.txt')

##############################################
# ファイルを削除する(File.delete)
##############################################

# File.delete('ファイルの置き場所')

File.delete('./assets/write_test2.txt')

# 同様に FileUtils.rm('ファイルの置き場所') でも削除できます

# FileUtils.rm('./assets/write_test2.txt')

##############################################
# ファイルの拡張子を調べる(File.extname)
##############################################

# File.extname('ファイルの置き場所')
# ただし仕様として最後の.以降を抜き出すようにできているようです。

p File.extname('./assets/file1.txt')
# => ".txt"

p File.extname('./assets/hoge.tar.gz')
# => ".gz"

######################################################
# ファイルがディレクトリかどうかを確認する(File.directory?)
######################################################

# File.directory?('ディレクトリの置き場所')
# ディレクトリの場合trueをそうでない場合はfalseを返す
# 「注意」ディレクトリが存在しない場合でもfalseを返す

# ディレクトリが存在する場合
p File.directory?('./assets')
# => true

# ディレクトリではない場合
p File.directory?('./assets/file1.txt')
# => false

# ディレクトリが存在しない場合
p File.directory?'./huuu'
# => false

##############################################
# ファイルやディレクトリの作成時刻を取得する(File.birthtime)
##############################################

# File.birthtime('ファイルもしくはディレクトリの場所')

# birthtimeからの返却はTimeクラスが返ります
p File.birthtime('./assets').class
# => Time

# ファイルの作成時刻を取得する場合
p File.birthtime('./assets/file1.txt')
# => 2019-10-12 19:15:24 +0900

# ディレクトリの作成時刻を取得する場合
p File.birthtime('./assets')
# => 2019-10-12 19:14:29 +0900

# ディレクトリが存在しない場合
begin
  p File.birthtime('./huuu')
rescue => e
  p e
  # => #<Errno::ENOENT: No such file or directory @ rb_file_s_birthtime - ./huuu>
end
