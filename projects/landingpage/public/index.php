<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <title>PIT – Local Development Pit Crew</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta name="description" content="PIT is a next-generation local development engine. Multi-PHP, cross-platform, and API-driven orchestration that makes your machine feel like a professional pit lane." />
  <script src="https://cdn.tailwindcss.com"></script>

  <style>
    :root {
      --pit-bg: #050608;
      --pit-surface: #101119;
      --pit-surface-soft: #151622;
      --pit-red: #e10600;
      --pit-red-soft: #ff4a3f;
      --pit-border: #262738;
      --pit-text-main: #f5f5f7;
      --pit-text-muted: #9ca3af;
      --pit-green: #00d68f;
      --pit-yellow: #ffd166;
    }

    * {
      box-sizing: border-box;
    }

    body {
      margin: 0;
      background: radial-gradient(circle at top, #111322 0, #050608 50%);
      color: var(--pit-text-main);
      font-family: system-ui, -apple-system, BlinkMacSystemFont, "Inter", sans-serif;
    }

    .pit-container {
      max-width: 1200px;
      margin: 0 auto;
      padding: 0 1.5rem;
    }

    .pit-glass {
      background: linear-gradient(145deg, rgba(255, 255, 255, 0.04), rgba(255, 255, 255, 0.02));
      border: 1px solid var(--pit-border);
      border-radius: 1.25rem;
      box-shadow:
        0 22px 45px rgba(0, 0, 0, 0.65),
        0 0 1px rgba(255, 255, 255, 0.04);
      backdrop-filter: blur(18px);
    }

    .pit-pill {
      border-radius: 999px;
      padding: 0.25rem 0.85rem;
      border: 1px solid rgba(255, 255, 255, 0.06);
      background: radial-gradient(circle at top left, rgba(225, 6, 0, 0.4), transparent 55%);
      font-size: 0.75rem;
      letter-spacing: 0.08em;
      text-transform: uppercase;
      color: var(--pit-text-muted);
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
    }

    .pit-pill-dot {
      width: 8px;
      height: 8px;
      border-radius: 999px;
      background: var(--pit-red);
      box-shadow: 0 0 14px rgba(225, 6, 0, 0.85);
    }

    .pit-btn-primary {
      background: linear-gradient(135deg, var(--pit-red) 0%, var(--pit-red-soft) 100%);
      color: white;
      border-radius: 999px;
      padding: 0.85rem 1.9rem;
      font-weight: 600;
      font-size: 0.95rem;
      border: none;
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
      box-shadow:
        0 14px 30px rgba(0, 0, 0, 0.75),
        0 0 10px rgba(225, 6, 0, 0.6);
      cursor: pointer;
      transition:
        transform 0.18s ease,
        box-shadow 0.18s ease,
        filter 0.18s ease;
    }

    .pit-btn-primary:hover {
      transform: translateY(-1px);
      filter: brightness(1.05);
      box-shadow:
        0 18px 40px rgba(0, 0, 0, 0.85),
        0 0 12px rgba(255, 74, 63, 0.8);
    }

    .pit-btn-ghost {
      border-radius: 999px;
      padding: 0.85rem 1.8rem;
      border: 1px solid rgba(148, 163, 184, 0.7);
      background: rgba(15, 23, 42, 0.4);
      color: var(--pit-text-main);
      font-size: 0.92rem;
      font-weight: 500;
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
      cursor: pointer;
      transition:
        background 0.18s ease,
        border-color 0.18s ease,
        transform 0.18s ease;
    }

    .pit-btn-ghost:hover {
      background: rgba(30, 64, 175, 0.4);
      border-color: rgba(248, 250, 252, 0.8);
      transform: translateY(-1px);
    }

    .pit-hero-grid {
      display: grid;
      gap: 3rem;
      grid-template-columns: minmax(0, 1.3fr) minmax(0, 1fr);
      align-items: center;
    }

    @media (max-width: 900px) {
      .pit-hero-grid {
        grid-template-columns: minmax(0, 1fr);
      }
    }

    .pit-logo {
      display: inline-flex;
      align-items: center;
      gap: 0.65rem;
    }

    .pit-logo-mark {
      width: 34px;
      height: 34px;
      border-radius: 0.9rem;
      background: radial-gradient(circle at 30% 0%, #ffffff 0, #d1d5db 38%, #111827 90%);
      display: flex;
      align-items: center;
      justify-content: center;
      box-shadow:
        0 10px 25px rgba(0, 0, 0, 0.9),
        0 0 10px rgba(255, 255, 255, 0.35);
    }

    .pit-logo-mark-inner {
      width: 18px;
      height: 12px;
      border-radius: 12px;
      border: 2px solid #111827;
      border-left-width: 4px;
      position: relative;
      overflow: hidden;
    }

    .pit-logo-mark-inner::after {
      content: "";
      position: absolute;
      inset: 0;
      border-radius: inherit;
      background: linear-gradient(135deg, #f9fafb, #e5e7eb);
      transform: translateX(-25%);
    }

    .pit-logo-text {
      font-weight: 800;
      font-size: 1.05rem;
      letter-spacing: 0.22em;
      text-transform: uppercase;
    }

    .pit-hero-title {
      font-size: clamp(2.6rem, 3.7vw, 3.4rem);
      line-height: 1.05;
      font-weight: 800;
      letter-spacing: -0.03em;
      margin-bottom: 1.25rem;
    }

    .pit-hero-subtitle {
      font-size: 1.13rem;
      color: var(--pit-text-muted);
      max-width: 34rem;
      line-height: 1.7;
      margin-bottom: 1.75rem;
    }

    .pit-metrics {
      display: flex;
      flex-wrap: wrap;
      gap: 1.75rem;
      margin-top: 2.4rem;
    }

    .pit-metric-item h4 {
      font-size: 1rem;
      color: var(--pit-text-muted);
      margin-bottom: 0.25rem;
    }

    .pit-metric-item p {
      font-size: 1.25rem;
      font-weight: 700;
    }

    .pit-metric-highlight {
      color: var(--pit-green);
    }

    .pit-card-heading {
      font-size: 0.9rem;
      text-transform: uppercase;
      letter-spacing: 0.16em;
      color: var(--pit-text-muted);
      margin-bottom: 0.35rem;
    }

    .pit-card-title {
      font-size: 1.05rem;
      font-weight: 600;
      margin-bottom: 0.75rem;
    }

    .pit-badge {
      display: inline-flex;
      align-items: center;
      padding: 0.2rem 0.7rem;
      border-radius: 999px;
      border: 1px solid rgba(148, 163, 184, 0.6);
      font-size: 0.72rem;
      text-transform: uppercase;
      letter-spacing: 0.12em;
      color: var(--pit-text-muted);
      gap: 0.25rem;
    }

    .pit-badge-dot-green {
      width: 8px;
      height: 8px;
      border-radius: 999px;
      background: var(--pit-green);
      box-shadow: 0 0 10px rgba(0, 214, 143, 0.9);
    }

    .pit-badge-dot-red {
      width: 8px;
      height: 8px;
      border-radius: 999px;
      background: var(--pit-red);
    }

    .pit-section-title {
      font-size: 2rem;
      font-weight: 700;
      letter-spacing: -0.03em;
      margin-bottom: 0.6rem;
    }

    .pit-section-subtitle {
      color: var(--pit-text-muted);
      max-width: 32rem;
      font-size: 0.98rem;
      line-height: 1.7;
    }

    .pit-feature-grid {
      display: grid;
      grid-template-columns: repeat(3, minmax(0, 1fr));
      gap: 1.7rem;
    }

    @media (max-width: 1024px) {
      .pit-feature-grid {
        grid-template-columns: repeat(2, minmax(0, 1fr));
      }
    }

    @media (max-width: 768px) {
      .pit-feature-grid {
        grid-template-columns: minmax(0, 1fr);
      }
    }

    .pit-feature-item {
      padding: 1.25rem 1.25rem 1.4rem;
      border-radius: 1.1rem;
      background: linear-gradient(150deg, rgba(15, 23, 42, 0.9), rgba(15, 23, 42, 0.7));
      border: 1px solid rgba(55, 65, 81, 0.9);
      position: relative;
      overflow: hidden;
    }

    .pit-feature-icon {
      width: 36px;
      height: 36px;
      border-radius: 0.9rem;
      background: radial-gradient(circle at 30% 0%, rgba(255, 255, 255, 0.9), rgba(148, 163, 184, 0.65));
      display: flex;
      align-items: center;
      justify-content: center;
      margin-bottom: 0.85rem;
      color: #111827;
      font-weight: 800;
      font-size: 1.05rem;
    }

    .pit-feature-label {
      font-size: 0.8rem;
      text-transform: uppercase;
      letter-spacing: 0.13em;
      color: var(--pit-text-muted);
      margin-bottom: 0.35rem;
    }

    .pit-feature-title {
      font-size: 1.03rem;
      font-weight: 600;
      margin-bottom: 0.4rem;
    }

    .pit-feature-text {
      font-size: 0.93rem;
      color: var(--pit-text-muted);
      line-height: 1.6;
    }

    .pit-chip-row {
      display: flex;
      flex-wrap: wrap;
      gap: 0.4rem;
      margin-top: 0.7rem;
    }

    .pit-chip {
      border-radius: 999px;
      border: 1px solid rgba(148, 163, 184, 0.6);
      padding: 0.1rem 0.6rem;
      font-size: 0.7rem;
      text-transform: uppercase;
      letter-spacing: 0.16em;
      color: #d1d5db;
    }

    .pit-code-block {
      font-family: "JetBrains Mono", "Fira Code", monospace;
      font-size: 0.82rem;
      line-height: 1.6;
      background: radial-gradient(circle at top left, #111827, #020617);
      padding: 1rem 1.2rem;
      border-radius: 0.8rem;
      border: 1px solid #1f2937;
      position: relative;
      overflow: auto;
      max-height: 260px;
    }

    .pit-code-toolbar {
      position: absolute;
      right: 0.7rem;
      top: 0.55rem;
      display: flex;
      align-items: center;
      gap: 0.4rem;
      font-size: 0.7rem;
      color: #9ca3af;
    }

    .pit-code-toolbar button {
      border-radius: 999px;
      border: 1px solid rgba(148, 163, 184, 0.6);
      padding: 0.2rem 0.55rem;
      background: rgba(15, 23, 42, 0.9);
      cursor: pointer;
      font-size: 0.68rem;
      color: #e5e7eb;
    }

    .pit-download-tabs {
      display: inline-flex;
      padding: 0.25rem;
      background: rgba(15, 23, 42, 0.95);
      border-radius: 999px;
      border: 1px solid rgba(55, 65, 81, 0.9);
      gap: 0.3rem;
    }

    .pit-download-tab {
      padding: 0.4rem 0.9rem;
      border-radius: 999px;
      font-size: 0.78rem;
      cursor: pointer;
      border: 1px solid transparent;
      color: var(--pit-text-muted);
      background: transparent;
      transition:
        background 0.16s ease,
        color 0.16s ease,
        border-color 0.16s ease;
    }

    .pit-download-tab.active {
      background: rgba(239, 68, 68, 0.08);
      border-color: rgba(248, 250, 252, 0.9);
      color: #f9fafb;
    }

    .pit-compare-table {
      width: 100%;
      border-collapse: collapse;
      font-size: 0.8rem;
    }

    .pit-compare-table th,
    .pit-compare-table td {
      padding: 0.65rem 0.75rem;
      border-bottom: 1px solid #1f2937;
    }

    .pit-compare-table thead {
      background: rgba(15, 23, 42, 0.95);
      color: #9ca3af;
      text-transform: uppercase;
      letter-spacing: 0.11em;
      font-size: 0.7rem;
    }

    .pit-label-yes {
      color: var(--pit-green);
      font-weight: 600;
    }

    .pit-label-no {
      color: #f97373;
      font-weight: 500;
    }

    .pit-label-partial {
      color: var(--pit-yellow);
    }

    .pit-tagline {
      text-transform: uppercase;
      letter-spacing: 0.2em;
      font-size: 0.75rem;
      color: var(--pit-text-muted);
    }

    .pit-step-grid {
      display: grid;
      grid-template-columns: repeat(3, minmax(0, 1fr));
      gap: 1.6rem;
    }

    @media (max-width: 900px) {
      .pit-step-grid {
        grid-template-columns: minmax(0, 1fr);
      }
    }

    .pit-step-number {
      width: 28px;
      height: 28px;
      border-radius: 999px;
      border: 1px solid rgba(148, 163, 184, 0.7);
      display: inline-flex;
      align-items: center;
      justify-content: center;
      font-size: 0.78rem;
      margin-bottom: 0.3rem;
      color: #e5e7eb;
    }

    .pit-section {
      padding: 4.5rem 0;
    }

    .pit-section-border {
      border-top: 1px solid rgba(31, 41, 55, 0.95);
    }

    .pit-fade {
      opacity: 0;
      transform: translateY(16px);
      transition:
        opacity 0.65s ease,
        transform 0.65s ease;
    }

    .pit-fade.visible {
      opacity: 1;
      transform: translateY(0);
    }

    .pit-footer-link {
      color: var(--pit-text-muted);
      font-size: 0.8rem;
    }

    .pit-footer-link:hover {
      color: #e5e7eb;
    }

    .pit-faq-item {
      border-radius: 1rem;
      border: 1px solid #1f2937;
      background: rgba(15, 23, 42, 0.85);
      padding: 1rem 1.2rem;
      margin-bottom: 0.75rem;
      cursor: pointer;
    }

    .pit-faq-question {
      display: flex;
      justify-content: space-between;
      align-items: center;
      gap: 1rem;
    }

    .pit-faq-title {
      font-size: 0.95rem;
      font-weight: 500;
    }

    .pit-faq-answer {
      font-size: 0.86rem;
      color: var(--pit-text-muted);
      margin-top: 0.45rem;
      display: none;
    }

    .pit-faq-item.open .pit-faq-answer {
      display: block;
    }

    .pit-faq-toggle {
      width: 18px;
      height: 18px;
      border-radius: 999px;
      border: 1px solid rgba(148, 163, 184, 0.8);
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 0.68rem;
      color: #e5e7eb;
      flex-shrink: 0;
    }

    .pit-badge-os {
      padding: 0.2rem 0.55rem;
      border-radius: 999px;
      border: 1px solid rgba(148, 163, 184, 0.6);
      font-size: 0.7rem;
      color: #d1d5db;
    }
  </style>
</head>

<body>

  <!-- NAVBAR -->
  <header class="sticky top-0 z-40 border-b border-slate-800/90 bg-black/70 backdrop-blur">
    <div class="pit-container flex items-center justify-between py-3">
      <div class="pit-logo">
        <div class="pit-logo-mark">
          <div class="pit-logo-mark-inner"></div>
        </div>
        <div>
          <div class="pit-logo-text tracking-[0.28em]">PIT</div>
          <div class="text-[0.65rem] uppercase tracking-[0.18em] text-slate-400">
            Local Dev Engine
          </div>
        </div>
      </div>

      <nav class="hidden md:flex items-center gap-7 text-[0.85rem] text-slate-300">
        <a href="#why" class="hover:text-white">Why PIT</a>
        <a href="#features" class="hover:text-white">Features</a>
        <a href="#downloads" class="hover:text-white">Downloads</a>
        <a href="#architecture" class="hover:text-white">Architecture</a>
        <a href="#faq" class="hover:text-white">FAQ</a>
      </nav>

      <div class="hidden md:flex items-center gap-3">
        <button class="pit-btn-ghost text-[0.82rem]">
          View on GitHub
        </button>
        <button class="pit-btn-primary text-[0.82rem]">
          Download
        </button>
      </div>
    </div>
  </header>

  <!-- HERO -->
  <main>
    <section class="pit-section">
      <div class="pit-container pit-hero-grid pit-fade">
        <div>
          <div class="pit-pill mb-4">
            <span class="pit-pill-dot"></span>
            Next-generation local dev engine
          </div>
          <h1 class="pit-hero-title">
            Local development that feels like an F1 pit lane.
          </h1>
          <p class="pit-hero-subtitle">
            PIT orchestrates web servers, PHP runtimes, and project isolation so you never think about ports, vhosts, or
            PHP versions again. Start coding, PIT handles the rest in the background.
          </p>

          <div class="flex flex-wrap gap-3 mb-3">
            <button class="pit-btn-primary">
              Download for your OS
            </button>
            <button class="pit-btn-ghost">
              Read documentation
            </button>
          </div>

          <div class="text-xs text-slate-500">
            No containers. No manual configs. No hidden magic.
          </div>

          <div class="pit-metrics">
            <div class="pit-metric-item">
              <h4>Cold start</h4>
              <p><span class="pit-metric-highlight">&lt; 500 ms</span></p>
            </div>
            <div class="pit-metric-item">
              <h4>Projects handled</h4>
              <p>50+ per machine</p>
            </div>
            <div class="pit-metric-item">
              <h4>PHP runtimes</h4>
              <p>7.4 · 8.0 · 8.1 · 8.3</p>
            </div>
          </div>
        </div>

        <!-- Hero Right: Status / cards -->
        <div class="space-y-4">
          <div class="pit-glass p-4">
            <div class="flex items-center justify-between mb-2">
              <span class="pit-card-heading">Live status</span>
              <span class="pit-badge">
                <span class="pit-badge-dot-green"></span>
                Engine running
              </span>
            </div>
            <div class="grid grid-cols-2 gap-3 text-[0.8rem]">
              <div class="bg-slate-900/80 rounded-lg border border-slate-700/80 p-3">
                <div class="text-slate-400 text-[0.72rem] uppercase tracking-[0.18em] mb-1">
                  Nginx
                </div>
                <div class="flex items-baseline gap-1">
                  <div class="text-sm font-semibold">Active</div>
                  <span class="text-[0.7rem] text-emerald-400">listening on 80, 443</span>
                </div>
                <div class="mt-2 h-1.5 rounded-full bg-slate-800 overflow-hidden">
                  <div class="h-full w-10/12 bg-emerald-500/90"></div>
                </div>
              </div>

              <div class="bg-slate-900/80 rounded-lg border border-slate-700/80 p-3">
                <div class="text-slate-400 text-[0.72rem] uppercase tracking-[0.18em] mb-1">
                  PHP-FPM pools
                </div>
                <div class="flex items-baseline gap-1">
                  <div class="text-sm font-semibold">6 pools</div>
                  <span class="text-[0.7rem] text-sky-400">3 versions</span>
                </div>
                <div class="mt-2 h-1.5 rounded-full bg-slate-800 overflow-hidden">
                  <div class="h-full w-9/12 bg-sky-500/90"></div>
                </div>
              </div>
            </div>

            <div class="mt-4 text-xs text-slate-400">
              PIT automatically starts and stops runtimes based on which projects you are actively using.
            </div>
          </div>

          <div class="pit-glass p-4">
            <div class="flex items-center justify-between mb-2">
              <div>
                <div class="pit-card-heading">Sample project</div>
                <div class="pit-card-title">myapp.test</div>
              </div>
              <span class="pit-badge">
                <span class="pit-badge-dot-red"></span>
                Laravel · PHP 8.3
              </span>
            </div>
            <div class="pit-code-block">
              <div class="pit-code-toolbar">
                <button data-copy="#code-cli">Copy</button>
              </div>
              <pre id="code-cli">
$ pit project:list

  1  myapp         php-8.3  nginx:8080   myapp.test
  2  legacy-app    php-7.4  nginx:8081   legacy.test
  3  api-service   php-8.1  nginx:8082   api.local

$ pit project:create blog --php=8.3

  ✓ Created blog runtime
  ✓ Generated nginx vhost blog.test
  ✓ Started php-fpm pool pit-blog-8.3
              </pre>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- WHY PIT -->
    <section id="why" class="pit-section pit-section-border">
      <div class="pit-container pit-fade">
        <div class="flex flex-col md:flex-row md:items-end md:justify-between gap-6 mb-10">
          <div>
            <div class="pit-tagline mb-2">Why PIT</div>
            <h2 class="pit-section-title">Local development has fallen behind.</h2>
            <p class="pit-section-subtitle">
              Containers are heavy. Old stacks like XAMPP feel abandoned. Valet and Herd lock you into specific
              operating systems. PIT is designed as a modern engine that respects your machine and your workflow:
              fast, predictable, cross-platform.
            </p>
          </div>
          <div class="bg-slate-950/80 rounded-2xl border border-slate-800 p-4 max-w-sm text-sm text-slate-300">
            <div class="font-semibold mb-1">Design principle</div>
            <p class="text-slate-400">
              Your editor and your application are the only things you should be thinking about.
              PIT silently does the work of an entire pit crew: booting runtimes, assigning ports, wiring virtual hosts,
              and guarding consistency.
            </p>
          </div>
        </div>

        <div class="grid md:grid-cols-3 gap-5">
          <div class="pit-feature-item">
            <div class="pit-feature-label">Problem</div>
            <div class="pit-feature-title">Too many tools, too little cohesion.</div>
            <p class="pit-feature-text">
              Developers juggle Nginx configs, PHP installers, DNS hacks, and ad-hoc scripts.
              Every project feels like a new set of manual steps.
            </p>
          </div>
          <div class="pit-feature-item">
            <div class="pit-feature-label">Insight</div>
            <div class="pit-feature-title">You only need a single orchestrator.</div>
            <p class="pit-feature-text">
              All of these steps are deterministic. They can be automated, versioned, and expressed as a concise model
              of "project runtime" and "machine capabilities".
            </p>
          </div>
          <div class="pit-feature-item">
            <div class="pit-feature-label">Response</div>
            <div class="pit-feature-title">PIT as your local pit lane.</div>
            <p class="pit-feature-text">
              PIT treats your machine like a race car in a pit lane. Projects come in, PIT configures exactly what they
              need, and you leave with a tuned environment with no manual chores.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- FEATURES -->
    <section id="features" class="pit-section pit-section-border">
      <div class="pit-container pit-fade">
        <div class="flex flex-col md:flex-row md:items-end md:justify-between gap-6 mb-10">
          <div>
            <div class="pit-tagline mb-2">Core capabilities</div>
            <h2 class="pit-section-title">Features built for serious work.</h2>
            <p class="pit-section-subtitle">
              Every piece of PIT is designed to be predictable, observable, and extensible. No black boxes. Everything
              you see in the panel can be controlled from the CLI and from the local HTTP API.
            </p>
          </div>
          <div class="flex flex-wrap gap-2 text-[0.8rem] text-slate-400">
            <span class="pit-badge-os">Linux</span>
            <span class="pit-badge-os">Windows</span>
            <span class="pit-badge-os">PHP 7.4 → 8.3</span>
            <span class="pit-badge-os">Nginx · Apache</span>
            <span class="pit-badge-os">REST API</span>
          </div>
        </div>

        <div class="pit-feature-grid">

          <article class="pit-feature-item">
            <div class="pit-feature-icon">R</div>
            <div class="pit-feature-title">Per-project runtimes</div>
            <p class="pit-feature-text">
              Each project receives its own FPM pool, log set, and vhost configuration. Switching PHP versions becomes
              a matter of updating one field in the project config.
            </p>
            <div class="pit-chip-row">
              <span class="pit-chip">Isolated pools</span>
              <span class="pit-chip">Per-project logs</span>
              <span class="pit-chip">Runtime profiles</span>
            </div>
          </article>

          <article class="pit-feature-item">
            <div class="pit-feature-icon">D</div>
            <div class="pit-feature-title">Automatic domains</div>
            <p class="pit-feature-text">
              PIT wires <span class="text-slate-200">*.test</span> style domains for every project. You focus on
              applications, not on editing hosts files or remembering ports.
            </p>
            <div class="pit-chip-row">
              <span class="pit-chip">TLD mapping</span>
              <span class="pit-chip">No browser hacks</span>
            </div>
          </article>

          <article class="pit-feature-item">
            <div class="pit-feature-icon">S</div>
            <div class="pit-feature-title">System aware engine</div>
            <p class="pit-feature-text">
              PIT inspects your machine capabilities and tracks which PHP binaries are available,
              automatically linking projects to supported runtimes.
            </p>
            <div class="pit-chip-row">
              <span class="pit-chip">Runtime discovery</span>
              <span class="pit-chip">Health checks</span>
            </div>
          </article>

          <article class="pit-feature-item">
            <div class="pit-feature-icon">A</div>
            <div class="pit-feature-title">API-first design</div>
            <p class="pit-feature-text">
              Everything PIT does is exposed over a local HTTP API. The CLI and the Panel are just two clients.
              You can build your own integrations, editor extensions, or dashboards.
            </p>
            <div class="pit-chip-row">
              <span class="pit-chip">Local REST</span>
              <span class="pit-chip">CLI + Panel</span>
              <span class="pit-chip">Automation hooks</span>
            </div>
          </article>

          <article class="pit-feature-item">
            <div class="pit-feature-icon">L</div>
            <div class="pit-feature-title">Logs you can trust</div>
            <p class="pit-feature-text">
              Unified access to nginx, FPM, and PIT engine logs. Open the panel, pick a project, and see
              exactly what is happening across the stack.
            </p>
            <div class="pit-chip-row">
              <span class="pit-chip">Per project log view</span>
              <span class="pit-chip">Engine logs</span>
            </div>
          </article>

          <article class="pit-feature-item">
            <div class="pit-feature-icon">F</div>
            <div class="pit-feature-title">F1-inspired speed</div>
            <p class="pit-feature-text">
              Written in Go with careful process orchestration. PIT is designed to stay out of your way,
              consuming minimal resources when idle and scaling up quickly when you start work.
            </p>
            <div class="pit-chip-row">
              <span class="pit-chip">Go engine</span>
              <span class="pit-chip">Low overhead</span>
            </div>
          </article>

        </div>
      </div>
    </section>

    <!-- DOWNLOADS -->
    <section id="downloads" class="pit-section pit-section-border">
      <div class="pit-container pit-fade">
        <div class="flex flex-col md:flex-row md:items-end md:justify-between gap-6 mb-8">
          <div>
            <div class="pit-tagline mb-2">Install PIT</div>
            <h2 class="pit-section-title">Choose your environment.</h2>
            <p class="pit-section-subtitle">
              PIT is distributed as a single binary with an optional panel.
              No installers, no wizards, no driver packages. Download, extract, and start the engine.
            </p>
          </div>
          <div class="flex flex-col items-start md:items-end gap-3">
            <div class="pit-download-tabs">
              <button class="pit-download-tab active" data-os="linux">Linux</button>
              <button class="pit-download-tab" data-os="windows">Windows</button>
            </div>
            <div class="text-xs text-slate-500">
              macOS support is planned. Follow the roadmap below.
            </div>
          </div>
        </div>

        <div class="grid md:grid-cols-2 gap-6">
          <div class="pit-glass p-5" id="download-card">
            <div class="flex items-center justify-between mb-4">
              <div>
                <div class="text-xs uppercase tracking-[0.18em] text-slate-400 mb-1">
                  Download package
                </div>
                <div class="font-semibold text-[1.05rem]">PIT Engine</div>
              </div>
              <span class="pit-badge">
                <span class="pit-badge-dot-green"></span>
                Stable channel
              </span>
            </div>
            <ul class="list-disc list-inside text-sm text-slate-300 space-y-1 mb-4">
              <li>Self-contained binary</li>
              <li>Built-in HTTP API and CLI</li>
              <li>Configured for Nginx and PHP-FPM</li>
            </ul>
            <button class="pit-btn-primary text-[0.85rem]">
              Download for Linux
            </button>
          </div>

          <div class="pit-glass p-5">
            <div class="text-xs uppercase tracking-[0.18em] text-slate-400 mb-2">
              One-line install
            </div>
            <div class="pit-code-block" id="code-install-block">
              <div class="pit-code-toolbar">
                <button data-copy="#code-install">Copy</button>
              </div>
              <pre id="code-install">
# Linux
curl -fsSL https://get.pit.dev/install.sh | bash

# Windows (PowerShell)
iwr https://get.pit.dev/install.ps1 -UseBasicParsing | iex
              </pre>
            </div>
            <p class="text-xs text-slate-500 mt-3">
              PIT never touches system files without explicit confirmation. Installation paths and permissions are
              visible, auditable, and reversible.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- ARCHITECTURE -->
    <section id="architecture" class="pit-section pit-section-border">
      <div class="pit-container pit-fade">
        <div class="flex flex-col md:flex-row md:items-end md:justify-between gap-6 mb-10">
          <div>
            <div class="pit-tagline mb-2">Architecture</div>
            <h2 class="pit-section-title">Engineered like a control room.</h2>
            <p class="pit-section-subtitle">
              PIT is a small, single binary written in Go. It exposes a local REST API, supervises child processes for
              Nginx and PHP-FPM, and keeps a consistent model of projects on your machine.
            </p>
          </div>
          <div class="bg-slate-950/90 rounded-2xl border border-slate-800 p-4 text-xs text-slate-300 max-w-sm">
            <div class="font-semibold mb-1">Design goals</div>
            <ul class="space-y-1 text-slate-400">
              <li>Single binary, no external runtime.</li>
              <li>Deterministic boot sequence.</li>
              <li>Clear boundary between engine and workloads.</li>
            </ul>
          </div>
        </div>

        <div class="grid md:grid-cols-2 gap-7">
          <div class="pit-glass p-5">
            <pre class="pit-code-block" style="max-height: 320px;">
+-----------------------------------------------------+
|                     PIT Engine                      |
|-----------------------------------------------------|
|  - Project registry                                 |
|  - Port allocator                                   |
|  - Runtime planner                                  |
|  - Health monitor                                   |
|  - Local REST API (127.0.0.1:7070)                  |
+-------------------------+---------------------------+
                          |
          +---------------+----------------+
          |                                |
+---------v----------+          +----------v---------+
|   Web frontends    |          |   PHP runtimes    |
|  (nginx / apache)  |          |   (PHP-FPM pools) |
+---------+----------+          +----------+--------+
          |                                |
   HTTP traffic                   FastCGI / sockets
          |                                |
+---------v----------+          +----------v---------+
|   Applications     |          |  Logs / metrics   |
|  (Laravel, CI4,    |          |  per project      |
|   WordPress, etc.) |          +-------------------+
+--------------------+
            </pre>
          </div>

          <div class="space-y-5">
            <div class="pit-glass p-4">
              <div class="text-xs uppercase tracking-[0.18em] text-slate-400 mb-1">
                Example API call
              </div>
              <div class="pit-code-block" style="max-height: 210px;">
                <pre>
GET http://127.0.0.1:7070/projects

[
  {
    "name": "myapp",
    "php": "8.3",
    "domain": "myapp.test",
    "nginxPort": 8080,
    "status": "running"
  },
  {
    "name": "legacy",
    "php": "7.4",
    "domain": "legacy.test",
    "nginxPort": 8081,
    "status": "stopped"
  }
]
                </pre>
              </div>
              <p class="text-xs text-slate-500 mt-2">
                The same model powers the panel, the CLI, and any automation you build around PIT.
              </p>
            </div>

            <div class="pit-glass p-4">
              <div class="text-xs uppercase tracking-[0.18em] text-slate-400 mb-2">
                Project lifecycle
              </div>
              <div class="pit-step-grid text-sm text-slate-300">
                <div>
                  <div class="pit-step-number">1</div>
                  <div class="font-semibold mb-1">Detect</div>
                  <p class="text-slate-400 text-[0.85rem]">
                    PIT scans your workspace, finding projects that match known patterns: public directories,
                    composer files, framework signatures.
                  </p>
                </div>
                <div>
                  <div class="pit-step-number">2</div>
                  <div class="font-semibold mb-1">Plan</div>
                  <p class="text-slate-400 text-[0.85rem]">
                    For each project, PIT assigns PHP version, pool name, port allocation, and domain, using defaults
                    that you can override.
                  </p>
                </div>
                <div>
                  <div class="pit-step-number">3</div>
                  <div class="font-semibold mb-1">Run</div>
                  <p class="text-slate-400 text-[0.85rem]">
                    The engine starts or stops runtimes on demand. When your editor opens a project, PIT ensures its
                    environment is fully ready in milliseconds.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </section>

    <!-- COMPARISON -->
    <section class="pit-section pit-section-border">
      <div class="pit-container pit-fade">
        <div class="flex flex-col md:flex-row md:items-end md:justify-between gap-6 mb-10">
          <div>
            <div class="pit-tagline mb-2">Positioning</div>
            <h2 class="pit-section-title">How PIT compares to existing tools.</h2>
            <p class="pit-section-subtitle">
              PIT is not a container platform and not a traditional all-in-one stack. It sits in between:
              orchestration like DDEV, simplicity like Laragon, but with an architecture that can evolve with how you
              build software.
            </p>
          </div>
          <div class="text-xs text-slate-500 max-w-xs">
            This table is deliberately opinionated. PIT is designed with full-time developers in mind, not only
            quick demos.
          </div>
        </div>

        <div class="pit-glass p-5 overflow-x-auto">
          <table class="pit-compare-table min-w-[640px]">
            <thead>
              <tr>
                <th>Feature</th>
                <th>PIT</th>
                <th>Laragon</th>
                <th>Valet</th>
                <th>DDEV</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Cross-platform</td>
                <td class="pit-label-yes">Linux, Windows</td>
                <td>Windows only</td>
                <td>macOS, Linux</td>
                <td>Cross-platform (Docker)</td>
              </tr>
              <tr>
                <td>Per-project PHP version</td>
                <td class="pit-label-yes">Yes</td>
                <td class="pit-label-no">Global only</td>
                <td class="pit-label-partial">Possible</td>
                <td class="pit-label-yes">Yes</td>
              </tr>
              <tr>
                <td>Container dependency</td>
                <td class="pit-label-no">None</td>
                <td class="pit-label-no">None</td>
                <td class="pit-label-no">None</td>
                <td class="pit-label-yes">Required</td>
              </tr>
              <tr>
                <td>API-driven engine</td>
                <td class="pit-label-yes">Yes</td>
                <td class="pit-label-no">No</td>
                <td class="pit-label-no">No</td>
                <td class="pit-label-partial">CLI focused</td>
              </tr>
              <tr>
                <td>Resource footprint</td>
                <td class="pit-label-yes">Very small</td>
                <td>Small</td>
                <td>Small</td>
                <td>Heavy</td>
              </tr>
              <tr>
                <td>Project isolation</td>
                <td class="pit-label-yes">FPM pools</td>
                <td class="pit-label-partial">Partial</td>
                <td class="pit-label-partial">FPM + symlinks</td>
                <td class="pit-label-yes">Containers</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </section>

    <!-- ROADMAP + USE CASES -->
    <section class="pit-section pit-section-border">
      <div class="pit-container pit-fade">
        <div class="grid md:grid-cols-2 gap-7">
          <div>
            <div class="pit-tagline mb-2">Use cases</div>
            <h2 class="pit-section-title">Where PIT fits.</h2>
            <p class="pit-section-subtitle mb-6">
              PIT is designed for people who build web applications every day and want their machine to behave like a
              reliable, predictable environment.
            </p>

            <div class="space-y-4 text-sm text-slate-300">
              <div class="flex gap-3">
                <div class="mt-1 h-1.5 w-6 rounded-full bg-emerald-400"></div>
                <div>
                  <div class="font-semibold">Full-stack web developers</div>
                  <p class="text-slate-400 text-[0.86rem]">
                    Run multiple Laravel, Symfony, CodeIgniter, or custom PHP services side by side without worrying
                    about port conflicts or global PHP mess.
                  </p>
                </div>
              </div>

              <div class="flex gap-3">
                <div class="mt-1 h-1.5 w-6 rounded-full bg-sky-400"></div>
                <div>
                  <div class="font-semibold">Agencies and internal teams</div>
                  <p class="text-slate-400 text-[0.86rem]">
                    Standardize local environments across a team even when machines differ.
                    A single PIT profile can describe the expected stack.
                  </p>
                </div>
              </div>

              <div class="flex gap-3">
                <div class="mt-1 h-1.5 w-6 rounded-full bg-rose-400"></div>
                <div>
                  <div class="font-semibold">Educators and students</div>
                  <p class="text-slate-400 text-[0.86rem]">
                    Replace complex multi-step installation guides with a single binary and a clear overview of how the
                    stack works. Students can focus on building.
                  </p>
                </div>
              </div>
            </div>
          </div>

          <div>
            <div class="pit-tagline mb-2">Roadmap</div>
            <h2 class="pit-section-title">Built to grow with you.</h2>
            <p class="pit-section-subtitle mb-6">
              PIT is intentionally small in scope, but the engine is structured so new capabilities can slot in without
              breaking existing workflows.
            </p>

            <div class="space-y-3 text-sm text-slate-300">
              <div class="pit-glass p-3">
                <div class="flex items-center justify-between">
                  <div class="font-semibold text-[0.9rem]">macOS support</div>
                  <span class="text-[0.7rem] text-emerald-400 font-medium">In progress</span>
                </div>
                <p class="text-[0.82rem] text-slate-400 mt-1">
                  Native launchd integration and keychain-friendly certificate handling.
                </p>
              </div>

              <div class="pit-glass p-3">
                <div class="flex items-center justify-between">
                  <div class="font-semibold text-[0.9rem]">Certificate manager</div>
                  <span class="text-[0.7rem] text-slate-400">Planned</span>
                </div>
                <p class="text-[0.82rem] text-slate-400 mt-1">
                  Automatic local TLS certificates for secure development on HTTPS URLs without browser warnings.
                </p>
              </div>

              <div class="pit-glass p-3">
                <div class="flex items-center justify-between">
                  <div class="font-semibold text-[0.9rem]">Editor integrations</div>
                  <span class="text-[0.7rem] text-slate-400">Planned</span>
                </div>
                <p class="text-[0.82rem] text-slate-400 mt-1">
                  VS Code and JetBrains plugins powered by the PIT API to show runtime status directly in the editor.
                </p>
              </div>
            </div>
          </div>

        </div>
      </div>
    </section>

    <!-- FAQ -->
    <section id="faq" class="pit-section pit-section-border">
      <div class="pit-container pit-fade">
        <div class="flex flex-col md:flex-row md:items-start md:justify-between gap-6 mb-8">
          <div>
            <div class="pit-tagline mb-2">FAQ</div>
            <h2 class="pit-section-title">Questions developers actually ask.</h2>
            <p class="pit-section-subtitle">
              PIT is engineered to be transparent. If you do not know what it is doing, that is a design bug. These are
              some of the most common questions.
            </p>
          </div>
        </div>

        <div class="grid md:grid-cols-2 gap-6">
          <div>

            <div class="pit-faq-item">
              <div class="pit-faq-question">
                <div class="pit-faq-title">Is PIT a container replacement?</div>
                <div class="pit-faq-toggle">+</div>
              </div>
              <div class="pit-faq-answer">
                PIT is not trying to replace containers. It is focused on the local, single-machine development
                experience. Many teams combine PIT for local work with containers or Kubernetes for staging and
                production.
              </div>
            </div>

            <div class="pit-faq-item">
              <div class="pit-faq-question">
                <div class="pit-faq-title">Can I keep my existing Nginx configuration?</div>
                <div class="pit-faq-toggle">+</div>
              </div>
              <div class="pit-faq-answer">
                PIT manages its own Nginx instance with isolated configuration. It does not overwrite system-wide
                Nginx. If you already run Nginx, you can either migrate to PIT's instance or keep them separated on
                different ports.
              </div>
            </div>

            <div class="pit-faq-item">
              <div class="pit-faq-question">
                <div class="pit-faq-title">How does PIT decide which PHP version a project uses?</div>
                <div class="pit-faq-toggle">+</div>
              </div>
              <div class="pit-faq-answer">
                By default, PIT chooses the newest available runtime compatible with the project's configuration file.
                You can override this explicitly per project using either the CLI or a small config file in the
                repository.
              </div>
            </div>

          </div>
          <div>

            <div class="pit-faq-item">
              <div class="pit-faq-question">
                <div class="pit-faq-title">What happens if something crashes?</div>
                <div class="pit-faq-toggle">+</div>
              </div>
              <div class="pit-faq-answer">
                The engines monitors child processes. If Nginx or a PHP-FPM pool exits unexpectedly, PIT captures
                logs, restarts the runtime, and surfaces the incident in both the CLI and the Panel so you can inspect
                what went wrong.
              </div>
            </div>

            <div class="pit-faq-item">
              <div class="pit-faq-question">
                <div class="pit-faq-title">Does PIT modify system files?</div>
                <div class="pit-faq-toggle">+</div>
              </div>
              <div class="pit-faq-answer">
                PIT is intentionally conservative. It does not change global PHP installations or server config.
                All of its state stays in its own directory unless you explicitly enable features such as host file
                integration.
              </div>
            </div>

            <div class="pit-faq-item">
              <div class="pit-faq-question">
                <div class="pit-faq-title">Can I script PIT for automation?</div>
                <div class="pit-faq-toggle">+</div>
              </div>
              <div class="pit-faq-answer">
                Yes. The local REST API is first-class. Any action available in the CLI is implemented on top of that
                API, so you can call it from your own scripts or CI tools.
              </div>
            </div>

          </div>
        </div>
      </div>
    </section>

    <!-- CTA FINAL -->
    <section class="pit-section pit-section-border">
      <div class="pit-container pit-fade">
        <div class="pit-glass px-6 py-9 flex flex-col md:flex-row md:items-center md:justify-between gap-6">
          <div>
            <div class="pit-tagline mb-2">Start today</div>
            <h2 class="pit-section-title mb-2">Let PIT handle the pit lane.</h2>
            <p class="pit-section-subtitle">
              Install the engine, point it to your projects, and feel the difference of a machine that behaves like a
              professional development environment, not a collection of ad-hoc scripts.
            </p>
          </div>
          <div class="flex flex-col items-start md:items-end gap-3">
            <button class="pit-btn-primary">
              Download PIT
            </button>
            <button class="pit-btn-ghost text-[0.8rem]">
              View quickstart guide
            </button>
          </div>
        </div>
      </div>
    </section>
  </main>

  <!-- FOOTER -->
  <footer class="border-t border-slate-800 bg-black py-6">
    <div class="pit-container flex flex-col md:flex-row md:items-center md:justify-between gap-4 text-slate-500 text-xs">
      <div>
        PIT is a local development engine designed with the same discipline as production infrastructure,
        but tuned for a developer’s day-to-day work.
      </div>
      <div class="flex flex-wrap gap-4">
        <a href="#" class="pit-footer-link">Documentation</a>
        <a href="#" class="pit-footer-link">GitHub</a>
        <a href="#" class="pit-footer-link">Roadmap</a>
        <a href="#" class="pit-footer-link">Security</a>
      </div>
    </div>
  </footer>

  <script>
    // Simple fade-in on scroll
    const fadeEls = document.querySelectorAll('.pit-fade');
    const observer = new IntersectionObserver(entries => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add('visible');
          observer.unobserve(entry.target);
        }
      });
    }, { threshold: 0.18 });

    fadeEls.forEach(el => observer.observe(el));

    // Copy buttons
    document.querySelectorAll('[data-copy]').forEach(btn => {
      btn.addEventListener('click', () => {
        const targetSelector = btn.getAttribute('data-copy');
        const el = document.querySelector(targetSelector);
        if (!el) return;
        const text = el.innerText;
        navigator.clipboard.writeText(text).then(() => {
          btn.textContent = 'Copied';
          setTimeout(() => { btn.textContent = 'Copy'; }, 1400);
        }).catch(() => {
          btn.textContent = 'Error';
          setTimeout(() => { btn.textContent = 'Copy'; }, 1400);
        });
      });
    });

    // FAQ toggle
    document.querySelectorAll('.pit-faq-item').forEach(item => {
      item.addEventListener('click', () => {
        item.classList.toggle('open');
      });
    });

    // Download tab switcher
    const tabs = document.querySelectorAll('.pit-download-tab');
    const downloadCard = document.getElementById('download-card');
    const updateDownloadCard = (os) => {
      const title = downloadCard.querySelector('.font-semibold');
      const list = downloadCard.querySelector('ul');
      const button = downloadCard.querySelector('button');

      if (os === 'linux') {
        title.textContent = 'PIT Engine for Linux';
        list.innerHTML = `
          <li>Static binary compiled for modern Linux distributions</li>
          <li>Systemd unit template included</li>
          <li>Works with your existing PHP installations or PIT-managed runtimes</li>
        `;
        button.textContent = 'Download for Linux';
      } else if (os === 'windows') {
        title.textContent = 'PIT Engine for Windows';
        list.innerHTML = `
          <li>Single executable, no installer required</li>
          <li>Optional service mode for automatic startup</li>
          <li>Integrated log viewer compatible with Windows paths</li>
        `;
        button.textContent = 'Download for Windows';
      }
    };

    tabs.forEach(tab => {
      tab.addEventListener('click', () => {
        tabs.forEach(t => t.classList.remove('active'));
        tab.classList.add('active');
        const os = tab.getAttribute('data-os');
        updateDownloadCard(os);
      });
    });
  </script>

</body>
</html>
