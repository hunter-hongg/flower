use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct PlanJson {
  pub workflow: String,
  pub steps: Vec<StepJson>,
}

#[derive(Debug, Deserialize)]
pub struct StepJson {
  pub name: String,
  pub exec: String,
  pub deps: Vec<String>,
}
