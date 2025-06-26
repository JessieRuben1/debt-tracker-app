# 💰 Debt Tracker

A full-stack web application for tracking personal debts, transactions, and monthly bills. Keep track of money you owe to others, money others owe you, and get insights into your monthly spending patterns through automated bank data integration.

## 🌟 Project Vision

This application aims to become a comprehensive personal finance management tool that combines:
- **Manual debt tracking** between friends and family
- **Automated bill tracking** through bank integration using the Kiwi API
- **Financial insights** and spending pattern analysis
- **Monthly financial overview** with automated categorization

## 🚀 Current Features

- ✅ User authentication (JWT-based)
- ✅ Contact management (add, edit, delete contacts)
- ✅ Multi-debt tracking per contact
- ✅ Transaction history and management
- ✅ Interactive analytics dashboard
- ✅ Debt settlement workflow
- ✅ Data export (CSV downloads)
- ✅ Real-time notifications
- ✅ Responsive web interface

## 🔮 Planned Features

- 🔄 **Bank Integration**: Kiwi API integration for automated bill tracking
- 📊 **Monthly Bill Analysis**: Categorized spending insights
- 📱 **Mobile Applications**: React Native/Flutter apps
- 🔔 **Smart Notifications**: Payment reminders and alerts
- 👥 **Group Expenses**: Split bills among multiple people
- 💱 **Multi-Currency**: Support for different currencies

## 🛠️ Tech Stack

- **Backend**: Go (Golang) + Gin + SQLite/PostgreSQL
- **Frontend**: HTML/CSS/JS + Alpine.js + Chart.js
- **Authentication**: JWT tokens
- **Future**: Kiwi API for bank data integration

## 📋 Prerequisites

⚠️ **Important**: This project is designed to run on **Linux/Ubuntu systems**. 

Windows users should use WSL (Windows Subsystem for Linux) with Ubuntu.

## 🚀 Quick Start

### For New Contributors

1. **Fork this repository** on GitHub
2. **Clone your fork**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/debt-tracker-app.git
   cd debt-tracker-app
   ```

3. **Follow the setup guide**: See [SETUP.md](./docs/SETUP.md) for detailed installation instructions

4. **Check project status**: See [PROGRESS.md](./docs/PROGRESS.md) for current development status

5. **Explore future plans**: See [ROADMAP.md](./docs/ROADMAP.md) for upcoming features

### For Project Collaborators

1. **Clone the main repository**:
   ```bash
   git clone https://github.com/MAIN_USERNAME/debt-tracker-app.git
   cd debt-tracker-app
   ```

2. **Create your feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Set up development environment**: Follow [SETUP.md](./docs/SETUP.md)

## 📚 Documentation

- **[SETUP.md](./docs/SETUP.md)** - Complete installation and setup guide
- **[PROGRESS.md](./docs/PROGRESS.md)** - Current development status and completed features  
- **[ROADMAP.md](./docs/ROADMAP.md)** - Future features and development phases
- **[API.md](./docs/API.md)** - API documentation and endpoints
- **[CONTRIBUTING.md](./docs/CONTRIBUTING.md)** - How to contribute to the project

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](./docs/CONTRIBUTING.md) for guidelines.

### Development Workflow

1. **Fork & Clone**: Get your own copy of the project
2. **Branch**: Create a feature branch for your work
3. **Develop**: Make your changes following our coding standards
4. **Test**: Ensure everything works on Ubuntu/Linux
5. **Pull Request**: Submit your changes for review

### Areas We Need Help With

- 🏦 **Bank API Integration**: Kiwi API implementation
- 📱 **Mobile Development**: React Native/Flutter apps  
- 🎨 **UI/UX Design**: Interface improvements
- 🧪 **Testing**: Unit and integration tests
- 📖 **Documentation**: Guides and tutorials
- 🐛 **Bug Fixes**: Issue resolution

## 🔒 Security & Privacy

This application handles sensitive financial data. We prioritize:
- Secure authentication and authorization
- Encrypted data storage
- Read-only bank access through Kiwi API
- No storage of bank credentials
- GDPR-compliant data handling

## 📞 Support

- **Issues**: Create GitHub issues for bugs and feature requests
- **Discussions**: Use GitHub Discussions for questions and ideas
- **Documentation**: Check the docs folder for detailed guides

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with Go and modern web technologies
- Kiwi API for secure bank data access
- Open-source community for excellent libraries

---

**Ready to manage your finances better? Let's build this together! 💪**