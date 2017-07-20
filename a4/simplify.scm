
(define simplify 
    (lambda (expr)
        (cond
            ((null? expr)
                '()
            )
            ((list? expr)
                (cond
                    ((equal? 2 (length expr))
                        (let 
                            (
                                (left (simplify(first expr)))
                                (right (simplify(second expr)))
                            )
                            (cond
                                ((and (equal? left 'inc) (number? right))
                                    (+ right 1)
                                )
                                ((and (equal? left 'dec) (number? right))
                                    (- right 1)
                                )
                                (else
                                    (list left right)
                                )
                            )
                        )
                    )
                    ((equal? 3 (length expr))
                        (let 
                            (
                                (left (simplify(first expr)))
                                (center (simplify(second expr)))
                                (right (simplify(third expr)))
                            )
                            (cond
                                ((and (equal? center '+) (equal? left 0))
                                    right
                                )
                                ((and (equal? center '+) (equal? right 0))
                                    left
                                )
                                ((and (equal? center '*) (equal? right 0))
                                    0
                                )
                                ((and (equal? center '*) (equal? left 0))
                                    0
                                )
                                ((and (equal? center '*) (equal? left 1))
                                    right
                                )
                                ((and (equal? center '*) (equal? right 1))
                                    left
                                )
                                ((and (equal? center '-) (equal? right 0))
                                    left
                                )
                                ((and (equal? center '-) (equal? left 0))
                                    right
                                )
                                ((and (equal? center '**) (equal? right 0))
                                    1
                                )
                                ((and (equal? center '**) (equal? right 1))
                                    left
                                )
                                ((and (equal? center '**) (equal? left 1))
                                    1
                                )
                                ((and (equal? center '-) (equal? left right))
                                    0
                                )
                                ((and (number? left) (number? right))
                                    ((operator center) left right)
                                )
                                (else
                                    (list left center right)
                                )
                            )
                        )
                    )
                    (else
                        (error "Unsupported expression.")
                    )
                )
            )
            (else
                expr
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