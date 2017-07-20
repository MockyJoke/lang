(define make-empty-env
    (lambda ()
        '()
    )
)


(define apply-env 
    (lambda (env v)
        (cond
             ((null? env) (error "apply-env: empty environment or variable not in environment"))
             ((equal? (car (car env)) v)
                (car (cdr (car env)))
             )
             (else (apply-env (cdr env) v))
        )
    )
)


(define extend-env 
    (lambda (v val env)
        (cond
            ((null? env)
                (list (list v val))
            )
            ((equal? (car (car env)) v)
                (cons (list v val) (cdr env))
            )
            (else
                (cons (car env) (extend-env v val(cdr env)))
            )
        )
    )
)
#|
(define t1
    (extend-env 'a 5
        (extend-env 'b 4
            (extend-env 'c 3
                (extend-env 'b 4
                    (make-empty-env)))))
)
|#
(define test-env
    (extend-env 'a 1
        (extend-env 'b 2
            (extend-env 'c 3
                (extend-env 'b 4
                    (make-empty-env)))))
)
