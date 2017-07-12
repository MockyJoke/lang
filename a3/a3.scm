;; Q1
(define my-last
    (lambda (lst)
        (cond
            ((null? lst) (error "my-last: empty list"))
            ((null? (cdr lst)) (car lst))
            (else (my-last (cdr lst)))
        )
    )
)

;; Q2
(define snoc
    (lambda (x lst)
        (append lst (list x))
    )
)

;; Q3
(define range
    (lambda (n)
        (cond
            ((<= n 0) '())
            (else (append (range (- n 1)) (list (- n 1))))
        )
    )
)

;; Q4
(define deep-sum
    (lambda (lst)
        (cond
            ((null? lst) 
                0
            )
            ((number? (car lst))
                (+ (car lst) (deep-sum (cdr lst)))
            )
            ((list? (car lst))
                (+ (deep-sum (car lst)) (deep-sum (cdr lst)))
            )
            (else
                (+ 0 (deep-sum (cdr lst)))
            )
        )
    )
)

;; Q5
(define count-primes
    (lambda (n)
        (cond
            ((< n 2) 0)
            (else
                (+ (if (is-prime? n) 1 0) (count-primes (- n 1)))
            ) 
        )
    )
)

(define is-prime?
    (lambda (n)
        (cond
            ((< n 2)
                #f
            )
            (else
                (is-prime-helper? n 2)
            )
        )
    )
)

;; Q6
(define is-prime-helper?
    (lambda (n i)
        (cond
            ((> i (sqrt n))
                #t
            )
            (else
                (and (not (= (modulo n i) 0)) (is-prime-helper? n (+ i 1)))
            )
        )
    )
)

;; Q7
(define is-bit?
    (lambda (x)
        (cond
            ((not (number? x))
                #f
            )
            ((or (= x 1) (= x 0))
                #t
            )
            (else
                #f
            )
        )
    )
)

(define is-bit-seq?
    (lambda (lst)
        (cond
            ((null? lst)
                #t
            )
            (else
                (and (is-bit? (car lst)) (is-bit-seq? (cdr lst)))
            )
        )
    )
)


;; Q8
(define all-bit-seqs
    (lambda (n)
        (cond
            ((< n 1)
                '()
            )
            ((= n 1)
                '((0) (1))
            )
            (else
                (append (append-all 0 (all-bit-seqs(- n 1))) (append-all 1 (all-bit-seqs(- n 1))))
            )
        )
    )
)

(define append-all
    (lambda (x lst)
        (cond
            ((null? lst)
                '()
            )
            (else
                (append (list (append (list x) (car lst))) (append-all x (cdr lst)))
            )
        )
    )
)