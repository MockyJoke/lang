(define make-empty-env
    (lambda ()
        '()
    )
)


(define apply-env 
    (lambda (env v)
        (cond
             ((null? env) (error "apply-env: empty environment"))
             ((equal? (car (car env)) v)
                (car (cdr (car env)))
             )
             (else (apply-env (cdr env) v))
        )
    )
)


(define has-env
    (lambda (env v)
        (cond
             ((null? env) #f)
             ((equal? (car (car env) v))
                #t
             )
             (else (has-env (cdr env) v))
        )
    )
)

(define extend-env 
    (lambda (v val env)
        (display "abc")
        (cond
            ((null? env)
                (begin (display "1")
                    (list (list v val))
                )
            )
            ((equal? (car (car env)) v)
                (begin (display "2")
                    (list (list v val) (cdr env))
                )
            )
            (else
                (begin (display "3")
                    (cons (car env) (extend-env v val(cdr env)))
                )
            )
        )
    )
)

;;(define k (extend-env 'a 1 (make-empty-env)))

#| 
(define test-env
    (extend-env 'a 1
        (extend-env 'b 2
            (extend-env 'c 3
                (extend-env 'b 4
                    (make-empty-env)))))
)
|#
