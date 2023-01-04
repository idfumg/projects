(ns a-01.conditionals)

(defn CondIf
  []
  (println "\nCondIf:")
  (if (= 5 5)
    (println "equal")
    (println "not equal")))

(CondIf)

(defn CondIfDo
  []
  (println "\nCondIfDoo")
  (if (= 5 5)
    (do (println "First do 1")
        (println "First do 2"))
    (do (println "Second do 1")
        (println "Second do 2"))))

(CondIfDo)

(defn CondNestedIf
  []
  (println "\nCondNestedIf")
  (if (and (= 5 5) (or (= 3 3) (not true)))
    (println "True")
    (println "False")))

(CondNestedIf)

(defn CondCase
  []
  (println "\nCondCase")
  (def pet "dog")
  (case pet
    "cat" (println "I have a cat")
    "dog" (println "I have a dog")
    "fish" (println "I have a fish")))

(CondCase)

(defn CondCond
  [amount]
  (println "\nCondCond")
  (cond
    (<= amount 2) (println "few")
    (<= amount 10) (println "several")
    (<= amount 100) (println "many")
    :else (println "loads")))

(CondCond 5)