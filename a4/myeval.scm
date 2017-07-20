
(load "env1.scm")

(define myeval
    (lambda (expr env)
        (cond
            ((null? expr)
                0
            )
            ((list? expr)
                (cond 
                    ((equal? 2 (length expr))
                        (cond
                            ((equal? (first expr) 'inc)
                                (+ (myeval (second expr) env) 1)
                            )
                            ((equal? (first expr) 'dec)
                                (- (myeval (second expr) env) 1)
                            )
                            (else
                                (error "Unsupported binary operator")
                            )
                        )
                    )
                    ((equal? 3 (length expr))
                        (cond
                            ((and (equal? '/ (second expr)) (equal? 0 (myeval (third expr) env)))
                                (error "Cannot divide by 0.")
                            )
                            (else
                                ((operator (second expr)) (myeval (first expr) env) (myeval (third expr) env))
                            )
                        )
                    )
                    (else
                        (error "Unsupported expression.")
                    )
                )
            )
            (else
                (cond
                    ((symbol? expr)
                        (apply-env env expr)
                    )
                    (else
                        expr
                    )
                )
            )
        )
    )
)

(define operator
    (lambda (e)
        (cond
            ((equal? e '+) +)
            ((equal? e '-) -)
            ((equal? e '*) *)
            ((equal? e '/) /)
            ((equal? e '**) expt)
            (else (error "Unsupported tenary operator"))
        )
    )
)

(define env1
    (extend-env 'x -1
        (extend-env 'y 4
            (extend-env 'x 1
                (make-empty-env))))
)

(define env2
    (extend-env 'm -1
        (extend-env 'a 4
            (make-empty-env)))
)

(define env3
    (extend-env 'q -1
        (extend-env 'r 4
            (make-empty-env)))
)
