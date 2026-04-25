import React, { useState, useEffect } from 'react';
import { Container, Typography, Box, Paper, Button, CircularProgress, Alert, useMediaQuery, useTheme } from '@mui/material';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';

interface FlowData {
  id: number;
  dir: string;
  success: number;
  failure: number;
}

const HomePage: React.FC = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  // const isTablet = useMediaQuery(theme.breakpoints.down('md'));
  
  const [flowData, setFlowData] = useState<FlowData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const fetchFlowData = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch('http://localhost:8080/api/flow');
      if (!response.ok) {
        throw new Error('Failed to fetch data');
      }
      const result = await response.json();
      setFlowData(result.data ?? []);
    } catch (err) {
      setError('Error fetching data. Please check if the backend is running.');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchFlowData();
  }, []);

  return (
    <Container maxWidth="lg" sx={{ py: { xs: 2, sm: 4 } }}>
      <Box 
        sx={{ 
          my: { xs: 2, sm: 4 }, 
          display: 'flex', 
          flexDirection: { xs: 'column', sm: 'row' },
          justifyContent: { xs: 'center', sm: 'space-between' }, 
          alignItems: { xs: 'center', sm: 'center' },
          gap: { xs: 2, sm: 0 },
          animation: 'fadeIn 0.5s ease-in-out'
        }}
      >
        <Typography 
          variant={isMobile ? "h5" : "h4"} 
          component="h1" 
          gutterBottom
          sx={{
            fontWeight: 700,
            color: '#64b5f6',
            marginBottom: { xs: '0.5rem', sm: '1rem' },
            textAlign: { xs: 'center', sm: 'left' }
          }}
        >
          Welcome to Flowd
        </Typography>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={fetchFlowData}
          disabled={loading}
          sx={{
            borderRadius: '8px',
            padding: { xs: '0.5rem 1rem', sm: '0.75rem 1.5rem' },
            fontWeight: 500,
            boxShadow: '0 4px 12px rgba(100, 181, 246, 0.3)',
            transition: 'all 0.3s ease',
            '&:hover': {
              transform: 'translateY(-2px)',
              boxShadow: '0 6px 16px rgba(100, 181, 246, 0.4)'
            },
            '&:disabled': {
              opacity: 0.6,
              transform: 'none',
              boxShadow: 'none'
            }
          }}
        >
          {loading ? <CircularProgress size={20} color="inherit" /> : 'Refresh'}
        </Button>
      </Box>
      
      {loading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', py: 10 }}>
          <CircularProgress size={isMobile ? 40 : 60} thickness={4} />
        </Box>
      ) : error ? (
        <Alert 
          severity="error" 
          sx={{ 
            mt: 4, 
            borderRadius: '8px',
            animation: 'fadeIn 0.5s ease-in-out'
          }}
        >
          {error}
        </Alert>
      ) : !flowData || flowData.length === 0 ? (
        <Paper 
          elevation={4} 
          sx={{ 
            p: { xs: 4, sm: 6 }, 
            mt: 4,
            borderRadius: '12px',
            background: 'rgba(255, 255, 255, 0.05)',
            backdropFilter: 'blur(10px)',
            animation: 'slideUp 0.5s ease-out',
            boxShadow: '0 8px 32px rgba(0, 0, 0, 0.15)',
            textAlign: 'center'
          }}
        >
          <Typography 
            variant={isMobile ? "h6" : "h5"} 
            sx={{
              fontWeight: 600,
              color: 'rgba(255, 255, 255, 0.7)',
              mb: 2
            }}
          >
            No Data Available
          </Typography>
          <Typography 
            variant="body1" 
            sx={{
              color: 'rgba(255, 255, 255, 0.5)',
              mb: 3
            }}
          >
            There is no flow data to display. Please check back later or ensure data is being collected.
          </Typography>
        </Paper>
      ) : (
        <Paper 
          elevation={4} 
          sx={{ 
            p: { xs: 2, sm: 4 }, 
            mt: 4,
            borderRadius: '12px',
            background: 'rgba(255, 255, 255, 0.05)',
            backdropFilter: 'blur(10px)',
            animation: 'slideUp 0.5s ease-out',
            boxShadow: '0 8px 32px rgba(0, 0, 0, 0.15)'
          }}
        >
          <Typography 
            variant={isMobile ? "h6" : "h5"} 
            gutterBottom
            sx={{
              fontWeight: 600,
              marginBottom: '1.5rem',
              color: 'rgba(255, 255, 255, 0.9)'
            }}
          >
            Flow Results
          </Typography>
          <Box sx={{ height: { xs: 350, sm: 500 }, width: '100%', minWidth: 0 }}>
            <ResponsiveContainer width="100%" height="100%">
              <BarChart
                data={flowData}
                margin={{ top: 20, right: 30, left: 20, bottom: isMobile ? 150 : 130 }}
              >
                <CartesianGrid strokeDasharray="3 3" stroke="rgba(255, 255, 255, 0.1)" />
                <XAxis
                  dataKey="dir"
                  tickFormatter={(dir) => dir.split('/').pop() || dir}
                  angle={isMobile ? -60 : -45}
                  textAnchor="end"
                  height={isMobile ? 140 : 120}
                  interval={0}
                />
                <YAxis 
                  allowDecimals={false} 
                  tick={{ fill: 'rgba(255, 255, 255, 0.7)' }}
                />
                <Tooltip
                  formatter={(value, name) => [`${value}`, name]}
                  labelFormatter={(label) => `Directory: ${label}`}
                  contentStyle={{
                    backgroundColor: 'rgba(26, 26, 46, 0.9)',
                    border: '1px solid rgba(100, 181, 246, 0.3)',
                    borderRadius: '8px',
                    color: 'rgba(255, 255, 255, 0.9)',
                    fontSize: isMobile ? '0.8rem' : '1rem'
                  }}
                />
                <Legend 
                  wrapperStyle={{
                    paddingTop: '1rem',
                    fontSize: isMobile ? '0.8rem' : '1rem'
                  }}
                  formatter={(value) => {
                    if (value === 'Success') {
                      return <span style={{ color: '#4caf50', fontWeight: '500' }}>{value}</span>;
                    } else if (value === 'Failure') {
                      return <span style={{ color: '#f44336', fontWeight: '500' }}>{value}</span>;
                    }
                    return <span style={{ color: 'rgba(255, 255, 255, 0.7)' }}>{value}</span>;
                  }}
                />
                <Bar 
                  dataKey="success" 
                  fill="#4caf50" 
                  name="Success" 
                  maxBarSize={isMobile ? 30 : 50}
                  radius={[4, 4, 0, 0]}
                  animationDuration={1500}
                />
                <Bar 
                  dataKey="failure" 
                  fill="#f44336" 
                  name="Failure" 
                  maxBarSize={isMobile ? 30 : 50}
                  radius={[4, 4, 0, 0]}
                  animationDuration={1500}
                />
              </BarChart>
            </ResponsiveContainer>
          </Box>
        </Paper>
      )}
      
      <Box sx={{ mt: 8, textAlign: 'center', animation: 'fadeIn 0.5s ease-in-out 0.3s both' }}>
        <Typography variant="body2" color="text.secondary">
          Flowd Dashboard • Real-time flow analysis
        </Typography>
      </Box>
      
      <style>{`
        @keyframes fadeIn {
          from {
            opacity: 0;
            transform: translateY(10px);
          }
          to {
            opacity: 1;
            transform: translateY(0);
          }
        }
        
        @keyframes slideUp {
          from {
            opacity: 0;
            transform: translateY(20px);
          }
          to {
            opacity: 1;
            transform: translateY(0);
          }
        }
        
        @media (max-width: 600px) {
          .MuiContainer-root {
            padding-left: 16px;
            padding-right: 16px;
          }
        }
      `}</style>
    </Container>
  );
};

export default HomePage;