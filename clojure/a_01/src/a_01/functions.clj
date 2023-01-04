(ns a_01.functions
    (:gen-class))

(defn -main
    "First function"
    []
    (println "My name is Artem")
    (println "loving Clojure so far")
    (+ 2 5))

(def increment (fn [x] (+ x 1)))

(defn increment_set
  [x]
  (map increment x))

(increment_set [1,2])

(defn DataTypes []
  (def a 1)
  (def b 1.125)
  (def c 1e9)
  (def d true)
  (def e "hahaha")
  (def f 'hello)
  
  (println a b c d e f))

(DataTypes)
