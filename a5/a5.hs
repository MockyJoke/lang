snoc :: a -> [a] -> [a]
snoc x [] = [x]
snoc x y = (head y) : snoc x (tail y)


myappend :: [a] -> [a] -> [a]
myappend x [] = x
myappend x y = myappend (snoc (head y) x) (tail y)


myreverse :: [a] -> [a]
myreverse [] = []
myreverse x = (last x):(myreverse (init x))


is_prime :: Int -> Bool
is_prime n 
    | n < 2 = False
    | otherwise = (length [x | x <- [2..n], n `mod` x == 0]) == 1

is_emirp :: Int -> Bool
is_emirp n = is_prime n && (read (myreverse (show n)) /= n) && is_prime (read (myreverse (show n)))

count_emirps :: Int -> Int
count_emirps n
    | n < 13 = 0
    | otherwise = length [x | x <- [13..n], is_emirp x]


biggest_sum :: [[Int]] -> [Int]
biggest_sum [x] = x 
biggest_sum lst = if sum (head lst) > sum (biggest_sum (tail lst)) then head lst else biggest_sum (tail lst)


greatest :: (a -> Int) -> [a] -> a
greatest f [x] = x 
greatest f lst = if f (head lst) > f (greatest f (tail lst) ) then head lst else greatest f (tail lst)


