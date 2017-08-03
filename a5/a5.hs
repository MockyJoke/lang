-- 1
snoc :: a -> [a] -> [a]
snoc x [] = [x]
snoc x y = (head y) : snoc x (tail y)


-- 2
myappend :: [a] -> [a] -> [a]
myappend x [] = x
myappend x y = myappend (snoc (head y) x) (tail y)


-- 3
myreverse :: [a] -> [a]
myreverse [] = []
myreverse x = (last x):(myreverse (init x))


-- 4
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


-- 5
biggest_sum :: [[Int]] -> [Int]
biggest_sum [x] = x 
biggest_sum lst = if sum (head lst) > sum (biggest_sum (tail lst)) then head lst else biggest_sum (tail lst)


-- 6
greatest :: (a -> Int) -> [a] -> a
greatest f [x] = x 
greatest f lst = if f (head lst) > f (greatest f (tail lst) ) then head lst else greatest f (tail lst)


-- 7
is_bit :: (Integral a) => a -> Bool
is_bit 1 = True
is_bit 0 = True
is_bit x = False


-- 8
flip_bit :: (Integral a) => a -> a
flip_bit x = if is_bit x then (if x == 0 then 1 else 0) else error "Input is not a bit."


-- 9
is_bit_seq1 :: (Integral a) => [a] -> Bool
is_bit_seq1 [] = True
is_bit_seq1 lst
    | is_bit (head lst) = True && is_bit_seq1 (tail lst)
    | otherwise = False

is_bit_seq2 :: (Integral a) => [a] -> Bool
is_bit_seq2 [] = True
is_bit_seq2 lst = if is_bit (head lst) then True && is_bit_seq2 (tail lst) else False

is_bit_seq3 :: (Integral a) => [a] -> Bool
is_bit_seq3 [] = True
is_bit_seq3 lst = all is_bit lst


-- 10
invert_bits1 :: (Integral a) => [a] -> [a]
invert_bits1 [] = []
invert_bits1 lst = flip_bit (head lst):invert_bits1(tail lst)

invert_bits2 :: (Integral a) => [a] -> [a]
invert_bits2 [] = []
invert_bits2 lst = map flip_bit lst

invert_bits3 :: (Integral a) => [a] -> [a]
invert_bits3 [] = []
invert_bits3 lst = [flip_bit x | x <- lst]


-- 11
bit_count :: (Integral a) => [a] -> (Int, Int)
bit_count [] = (0, 0)
bit_count lst = (length [x|x <- lst, x == 0], length [x|x <- lst, x == 1])


-- 12
all_basic_bit_seqs :: (Integral a) => a -> [[Int]]
all_basic_bit_seqs n
    | n < 1 = []
    | n == 1 = [[0],[1]]
    | otherwise = myappend [0:x | x <- all_basic_bit_seqs (n-1) ] [1:x | x <- all_basic_bit_seqs (n-1) ]



data Bit = Zero | One
    deriving (Show, Eq)

-- 13
flipBit :: Bit -> Bit
flipBit Zero = One
flipBit One = Zero


-- 14
invert :: [Bit] -> [Bit]
invert [] = []
invert lst = flipBit (head lst) : invert (tail lst)


-- 15
all_bit_seqs :: (Integral a) => a -> [[Bit]]
all_bit_seqs n
    | n < 1 = []
    -- I came up with this solution based on getting binary digit by continuous division.
    | otherwise = [ [ if ((x `quot` (2^q)) `mod` 2) == 0 then Zero else One | q<-[n-1,n-2..0]]  | x <- [0..2^n-1]]
    

-- 16
bitSum1 :: [Bit] -> Int
bitSum1 [] = 0
bitSum1 (x:xs) = (if x == One then 1 else 0) + bitSum1(xs)


-- 17
bitSum2 :: [Maybe Bit] -> Int
bitSum2 [] = 0
bitSum2 (Nothing:xs) = bitSum2 xs
bitSum2 ((Just x):xs) = (if x == One then 1 else 0) + bitSum2(xs)



data List a = Empty | Cons a (List a)
    deriving Show


-- 18
toList :: [a] -> List a
toList [] = Empty
toList (x:xs) = Cons x (toList xs)


-- 19
toHaskellList :: List a -> [a]
toHaskellList Empty = []
toHaskellList (Cons x (xs)) = x:(toHaskellList xs)


-- 20
append :: List a -> List a -> List a
append Empty y = y
append (Cons x (xs)) y = Cons x (append xs y)


-- 21
removeAll :: (a -> Bool) -> List a -> List a
removeAll f Empty = Empty
removeAll f (Cons x (xs)) = if (f x) then removeAll f xs else Cons x (removeAll f xs)


-- 22
-- I got the idea of implementing quick sort in haskell for list from
-- https://smthngsmwhr.wordpress.com/2012/11/09/sorting-algorithms-in-haskell/
-- Helper for fitering custom defined List
listFilter :: Ord a => (a -> a -> Bool) -> List a -> a -> List a
listFilter f Empty y = Empty
listFilter f (Cons x (xs)) y =  if (f x y) then Cons x (listFilter f xs y) else listFilter f xs y


sort :: Ord a => List a -> List a
sort Empty = Empty
sort (Cons x (xs)) = append (sort (listFilter (<=) xs x)) (Cons x (sort (listFilter (>) xs x)))

