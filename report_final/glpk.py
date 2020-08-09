import pyomo.environ as pyo
import sys

model = pyo.ConcreteModel(name="LP sample",
                          doc="4 variables, 3 constraints")

model.x1 = pyo.Var(domain=pyo.NonNegativeReals)
model.x2 = pyo.Var(domain=pyo.NonNegativeReals)
model.x3 = pyo.Var(domain=pyo.NonNegativeReals)
model.x4 = pyo.Var(domain=pyo.NonNegativeReals)

model.OBJ = pyo.Objective(expr = 200*model.x1 + 150*model.x2 + 100*model.x3 + 100*model.x4,
                          sense = pyo.maximize)

model.Constraint1 = pyo.Constraint(expr = 10*model.x1 + 5*model.x2 <= 1000)
model.Constraint2 = pyo.Constraint(expr = 3*model.x2 + 8*model.x3 + 2*model.x4 <= int(sys.argv[1]))
model.Constraint3 = pyo.Constraint(expr = 2*model.x2 + 2*model.x3 + 8*model.x4 <= 3000)

opt = pyo.SolverFactory('glpk')
res = opt.solve(model)

#benefit = float(f"{model.OBJ()}") - 10*float(f"{model.x1()}") - 7*float(f"{model.x2()}")
#print(f"{model.x1()}",sys.argv[1])
#print(f"{model.x2()}",sys.argv[1])
#print(f"{model.x3()}",sys.argv[1])
print(f"{model.x4()}",sys.argv[1])
#print(benefit,1000)
#print(f"{model.OBJ()}",1000)
