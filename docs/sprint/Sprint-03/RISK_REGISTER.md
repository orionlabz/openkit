# Sprint 03 Risk Register

| Risk | Impact | Likelihood | Mitigation |
| --- | --- | --- | --- |
| Scripts drift from template references | Medium | Medium | Keep scripts under `base/scripts/` and ensure references match `.opencode/scripts/` |
| Python not installed | Medium | Low | Document requirement; scripts fail clearly |
| Docker Compose absent | Low | Medium | Script prints clear error and exits non-zero |
