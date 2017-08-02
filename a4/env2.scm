(define make-empty-env
    (lambda ()
        (lambda (v)
            (error "apply-env: empty environment")
        )
    )
)

(define apply-env 
    (lambda (env v) 
        (env v) ; env returns the value of v if v is bound
    )
)

(define extend-env 
    (lambda (v val env)
        (lambda (x)
            (cond
                ((equal? x v)
                    val
                )
                (else
                    (env x)
                )
            )
        )
    )
)

(define test-env
    (extend-env 'a 1
        (extend-env 'b 2
            (extend-env 'c 3
                (extend-env 'b 4
                    (make-empty-env)))))
)
