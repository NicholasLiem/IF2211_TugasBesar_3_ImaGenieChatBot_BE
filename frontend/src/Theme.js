import { extendTheme } from '@chakra-ui/react';
import Button from './styles/ButtonStyles';

const customTheme = extendTheme({
  components: {
    Button,
  },
});

export default customTheme;