##############################################
# case文を使った条件分岐の書き方
##############################################

# if文の条件分岐と異なり複数の条件の際の分岐に向いています

# 1の時 lucky
# 2~4の時 unlucky
# それ以外の時 unknown
# を返します
# 下記のようにメソッドを使う場合はreturnで返してしまえばいいです。
# enumなどとよく組み合わせて使われます

def lucky(input)
  case input
    when 1
      return 'lucky'
    when 2..4
      return 'unlucky'
    else
      return 'unknown'
  end
end

p lucky(1)
# => "lucky"

p lucky(2)
# => "unlucky"

p lucky(4)
# => "unlucky"

p lucky(5)
# => "unknown"
