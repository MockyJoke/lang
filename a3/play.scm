(define hello
    (lambda (name)
        (string-append "Hello " name "!")
    )
)
(define sum
    (lambda (a b c)
        (+ a b c)
    )
)

(define remove-all
    (lambda (x lst)
        (cond
            ((null? lst)
                lst
            )

            ((equal? x (car lst))
                (remove-all x (cdr lst))
            )

            (else
                (cons (car lst) 
                      (remove-all x (cdr lst))
                )
            ) 
        )
    )
)

(define add1
    (lambda (num)
        (+ 1 num)
    )
)

(define minus1
    (lambda (num)
        (- 1 num)
    )
)

